# MongoDB Atlas Providers

Building the extension provider can be done with the following command (in each cloud directory):

```bash
make install
```

This will build the runtime provider and the deployment provider, packaging them together and saving it to `$HOME/.nitric/providers/mongo/{cloud-name}-0.0.1`.

To use the custom extension you can use the following stack configuration file. It requires you fill in the orgid to deploy your cluster to MongoDB Atlas.

```yaml
provider: custom/extension@0.0.1
region: us-east-1
orgid: xxxxxxxx
```

When using `nitric up` or `nitric down` you will need to have the following environment variables set. These will both be kept secret when deploying using Pulumi.

- MONGODB_ATLAS_PUBLIC_KEY
- MONGODB_ATLAS_PRIVATE_KEY
