# MongoDB::Atlas::CloudBackupSnapshot

## Description
Returns, takes, edits, and removes Cloud Backup snapshots.

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/cloud-backup-snapshot
./test/cloud-backup-snapshot.create-sample-cfn-request.sh YourProjectID ClusterName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.

## Installation
TAGS=logging make
cfn submit --verbose --set-default

## CloudFormation Examples

Please see the [CFN Template](/examples/cloud-backup-snapshot/snapshot.json) for example resource

## Integration Testing w/ AWS

The [/quickstart-mongodb-atlas/scripts/launch-quickstart.sh](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/quickstart-mongodb-atlas/scripts/launch-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other necessary parameters.

You can use the cloud-backup-snapshot.sample-template.yaml to create a stack using the resource.
Similar to the local testing described above you can follow the logs for the deployed
lambda function which handles the request for the Resource Type.

In one shell session:
```
aws logs tail mongodb-atlas-backup-snapshot-logs --follow
```

And then you can create the stack with a helper script it insert the apikeys for you:


```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-quickstart.sh ${repo_root}/cfn-resources/cloud-backup-snapshot/test/cloud-backup-snapshot.sample-template.yaml SampleAccessList1 ParameterKey=ProjectId,ParameterValue=<YOUR_PROJECT_ID> ParameterKey=ClusterName,ParameterValue=<ClusterName> ParameterKey=RetentionInDays,ParameterValue=<RetentionInDays> ParameterKey=Description,ParameterValue=<Description> 
```

For more information see: MongoDB Atlas API [Cloud Backups](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Cloud-Backups) Documentation.
