# jugaduInTerratest

https://www.linkedin.com/pulse/jugadu-terratest-iac-testing-riju-bhattacharya

![Alt text](./terratestWorkflow.jpg?raw=true "Terratest Workflow")

To test with the script follow the steps below:

# 1. Clone this repository in Azure CloudShell [ with Permission to create Azure StorageAccount ]
  ~] git clone https://github.com/rijub2019/jugaduInTerratest.git

# 2. Change directory to 'jugaduInTerratest'  
  ~] cd jugaduInTerratest
  
# 3. Download required go packages
  ~] go mod azure_storage_test.go \
  ~] go mod tidy
  
# 4. Execute the test
  ~] go test -v azure_storage_test.go
  
*Addiotnally you can edit 'terraform/variables.tf' line no.35 to change the 'default' value for enabling versioning to 'false', to check correct functioning of the script.\
*You will find terraform, go, git already installed in Azure CloudShell.

