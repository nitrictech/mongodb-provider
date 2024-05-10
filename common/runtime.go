package common

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	grpc_errors "github.com/nitrictech/nitric/core/pkg/grpc/errors"
	kvstorepb "github.com/nitrictech/nitric/core/pkg/proto/kvstore/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

type MongoDBServer struct {
	client *mongo.Client
}

var _ kvstorepb.KvStoreServer = &MongoDBServer{}

func (k *MongoDBServer) getCollectionHandle(collection string) *mongo.Collection {
	return k.client.Database("nitric").Collection(collection)
}

// Updates a secret, creating a new one if it doesn't already exist
// Get an existing document
func (k *MongoDBServer) GetValue(ctx context.Context, req *kvstorepb.KvStoreGetValueRequest) (*kvstorepb.KvStoreGetValueResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("MongoDBServer.GetValue")

	coll := k.getCollectionHandle(req.Ref.Store)

	filter := bson.D{{"_id", req.Ref.Key}}

	res := coll.FindOne(ctx, filter)
	if res == nil {
		return nil, newErr(
			codes.NotFound,
			fmt.Sprintf("key %s not found in store %s", req.Ref.Key, req.Ref.Store),
			fmt.Errorf(""),
		)
	}

	var result primitive.M
	err := res.Decode(&result)
	if err != nil {
		return nil, newErr(
			codes.Internal,
			"unable to convert value to raw bson",
			err,
		)
	}

	b, err := json.Marshal(result)
	if err != nil {
		return nil, newErr(
			codes.Internal,
			"unable to convert value to raw json",
			err,
		)
	}

	var structContent structpb.Struct
	err = proto.Unmarshal(b, &structContent)
	if err != nil {
		return nil, newErr(
			codes.Internal,
			"unable to convert value to pb struct",
			err,
		)
	}

	return &kvstorepb.KvStoreGetValueResponse{
		Value: &kvstorepb.Value{
			Ref:     req.Ref,
			Content: &structContent,
		},
	}, nil
}

// Create a new or overwrite an existing document
func (k *MongoDBServer) SetValue(ctx context.Context, req *kvstorepb.KvStoreSetValueRequest) (*kvstorepb.KvStoreSetValueResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("MongoDBServer.SetValue")

	coll := k.getCollectionHandle(req.Ref.Store)

	contents := req.Content.AsMap()
	contents["_id"] = req.Ref.Key

	_, err := coll.InsertOne(ctx, contents)
	if err != nil {
		return nil, newErr(
			codes.Internal,
			fmt.Sprintf("unable to insert %s into %s store", req.Ref.Key, req.Ref.Store),
			err,
		)
	}

	return &kvstorepb.KvStoreSetValueResponse{}, nil
}

// Delete an existing document
func (k *MongoDBServer) DeleteKey(ctx context.Context, req *kvstorepb.KvStoreDeleteKeyRequest) (*kvstorepb.KvStoreDeleteKeyResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("MongoDBServer.DeleteValue")

	coll := k.getCollectionHandle(req.Ref.Store)

	filter := bson.D{{"_id", req.Ref.Key}}

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, newErr(
			codes.Internal,
			fmt.Sprintf("unable to delete %s from %s store", req.Ref.Key, req.Ref.Store),
			err,
		)
	}

	return &kvstorepb.KvStoreDeleteKeyResponse{}, nil
}

// Iterate over all keys in a store
func (k *MongoDBServer) ScanKeys(req *kvstorepb.KvStoreScanKeysRequest, stream kvstorepb.KvStore_ScanKeysServer) error {
	newErr := grpc_errors.ErrorsWithScope("MongoDBServer.ScanKeys")

	coll := k.getCollectionHandle(req.Store.Name)

	regex := primitive.Regex{Pattern: "^" + req.Prefix, Options: ""}

	// Define your aggregation pipeline
	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{{"_id", bson.D{{"$regex", regex}}}}},
			// Add other stages as needed
		},
	}

	// Perform the aggregation
	cursor, err := coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		return newErr(
			codes.Internal,
			fmt.Sprintf("unable to scan keys with prefix %s from %s store", req.Prefix, req.Store.Name),
			err,
		)
	}
	defer cursor.Close(context.Background())

	// Iterate over the results
	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		key := result["_id"].(string)

		if err := stream.Send(&kvstorepb.KvStoreScanKeysResponse{
			Key: key,
		}); err != nil {
			return newErr(
				codes.Internal,
				"failed to send response",
				err,
			)
		}
	}

	return nil
}

func New() (*MongoDBServer, error) {
	ctx := context.TODO()

	url := os.Getenv("MONGO_CLUSTER_CONNECTION_STRING")
	if url == "" {
		return nil, fmt.Errorf("MONGO_CLUSTER_CONNECTION_STRING is unset")
	}

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, err
	}

	return &MongoDBServer{
		client: client,
	}, nil
}
