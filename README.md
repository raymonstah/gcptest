# GCP Test

Just wanting to get a feel of how fast GCP Cloud Functions are.
This is testing making an HTTP request with a simple lookup in Firestore.


### Some notes:

Set credentials for Firestore stuff:
```shell script
export GOOGLE_APPLICATION_CREDENTIALS="gcptest-BLAH.json"

```

Deploy a function: 
 ```shell script
gcloud functions deploy GetHelloFirestore --runtime go113 --trigger-http --allow-unauthenticated --region=us-west2 --memory=128MB
```