# Debug and Run Tests Using Go

  1. Download your level state file from azure blob storage.
  2. Rename the file to terraform.tfstate and place it in a known location.
  3. Create a .env file in the tests folder with the following structure.

      ```env
      STATE_FILE_PATH=<path to state file>
      ENVIRONMENT=<deployed caf environment name>
      ARM_SUBSCRIPTION_ID=<subscription id>
      ```

  4. Run the following command in the ./tests folder to download go dependencies.

      ```go
      go get
      go mod tidy
      ```

  5. Run your test using go test cmd

      ```go
      // Run tests
      go test
      ```
