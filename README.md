# gogpg
Simplified GPG encryption using golang

## Generating GPG keys

1. List all available keys

    ```
    gpg --list-keys
    ```

2. Create a new key

    The adviced configuration is:

    | Question    | Answer                                |
    | ---         | ---                                   |
    | kind of key | (1) RSA and RSA (default)             |
    | keysize     | 4096                                  |
    | *validity   | 0 = key does not expire (for staging) |

    *validity should be specified for production.

    ```
    gpg --full-generate-key
    ```

3. Export the keys

    ```
    gpg --armor --output <uid>.gpg.pub --export <uid>
    gpg --armor --output <uid>.gpg.pvt --export-secret-keys <uid>
    ```

