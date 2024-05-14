import mongoose, { Schema } from "mongoose";

const uri = process.env.MONGO_CLUSTER_CONNECTION_STRING;

const Profile = mongoose.model(
  "Profile",
  new Schema({ _id: String, name: String })
);

async function connect() {
  try {
    // Create a Mongoose client with a MongoClientOptions object to set the Stable API version
    await mongoose.connect(uri);

    await mongoose.connection.db.admin().command({ ping: 1 });

    console.log(
      "Pinged your deployment. You successfully connected to MongoDB!"
    );
  } catch {
    // Ensures that the client will close when you error
    console.log("Ping failed. Connection unsuccessful.");
    await mongoose.disconnect();
  }
}

export { connect, Profile };
