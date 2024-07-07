# Basic Example

This basic example demonstrates how to use the Logto API client to get a single organization by ID and then list all of its users. I assume that you are running Logto locally in Docker and it is available at `http://localhost:3002`.

1. Follow this [guide](https://docs.logto.io/docs/recipes/interact-with-management-api/) to set up an M2M app and its permissions and get the client ID and client secret.
2. Add an organization and add some members (that you already created in User management), and then copy the organization ID.
2. Run the `get-token.sh` script to get an access token, which will be saved in a file named `token.json`. Follow the instructions in the script to set the `CLIENT_ID` and `CLIENT_SECRET` environment variables.

    ```bash
    ./get-token.sh
    ```

3. Run the `main.go` file to get an organization by ID and list all of its users. Consider replacing the organization ID with the one you copied in step 2. The output will be similar to the following:

    ```bash
    go run main.go sxufs5cz6iy5
    (*logto.ListApplicationUserConsentOrganizations200ResponseOrganizationsInner)(0x...)({
        TenantId: (string) (len=7) "default",
        Id: (string) (len=12) "sxufs5cz6iy5",
        Name: (string) (len=4) "test",
        Description: (logto.NullableString) {
        value: (*string)(0x...)((len=4) "Test"),
        isSet: (bool) true
        },
        ...
    })
    ([]logto.ListOrganizationUsers200ResponseInner) (len=1 cap=1) {
        (logto.ListOrganizationUsers200ResponseInner) {
            Id: (string) (len=12) "lz3u1tgjzkm1",
            Username: (logto.NullableString) {
            value: (*string)(0x...)((len=4) "test"),
            isSet: (bool) true
            },
            PrimaryEmail: (logto.NullableString) {
            value: (*string)(0x...)((len=16) "test@example.com"),
            isSet: (bool) true
        },
        ...
        OrganizationRoles: ([]logto.ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) {}
        }
    }
    ```
