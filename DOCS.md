Use this plugin to upload files and build artifacts to an Google Cloud Storage
bucket.

## Config

The following parameters are used to configure the plugin:

* **credentials** - json contents of the google application credentials (see https://developers.google.com/identity/protocols/application-default-credentials)
* **bucket** - bucket name
* **acl** - access to files that are uploaded (`private`, `public`)
* **source** - source location of the files, using a glob matching pattern
* **target** - target location of files in the bucket
* **strip_prefix** - strip the prefix from source path
* **exclude** - glob exclusion patterns
* **compress** - gzip files before they are uploaded

The following secret values can be set to configure the plugin.

* **GOOGLE_APPLICATION_CREDENTIALS_CONTENTS** - corresponds to **webhook**
* **GCS_BUCKET** - corresponds to **webhook**

It is highly recommended to put the **GOOGLE_APPLICATION_CREDENTIALS_CONTENTS**
into a secret so it is not exposed to users. This can be done using the
drone-cli.

```bash
drone secret add --image=wyattjoh/drone-gcs \
    octocat/hello-world GOOGLE_APPLICATION_CREDENTIALS_CONTENTS @/path/to/application_credentials.json
```

Then sign the YAML file after all secrets are added.

```bash
drone sign octocat/hello-world
```

See [secrets](http://readme.drone.io/0.5/usage/secrets/) for additional
information on secrets

## Example

Common example to upload to gcs:

```yaml
pipeline:
  gcs:
    image: wyattjoh/drone-gcs
    acl: public
    bucket: "my-bucket-name"
    credentials: ${GOOGLE_APPLICATION_CREDENTIALS_CONTENTS}
    source: public/**/*
    strip_prefix: public/
    target: /target/location
    exclude:
      - **/*.xml
```
