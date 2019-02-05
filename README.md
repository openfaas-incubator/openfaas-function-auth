openfaas-function-auth
==============

Examples of authentication in OpenFaaS Serverless functions.

## 1. Basic auth

Username password protects governs access for your function.

* Use secrets or environmental variables for credentials

```bash
$ faas-cli template store pull golang-http

$ faas-cli secret create fn-basic-auth-username --from-literal="admin"

$ export PASSWORD=$(head -c 16 /dev/urandom | shasum | cut -d" " -f1)
$ echo $PASSWORD
$ echo -n $PASSWORD | faas-cli secret create fn-basic-auth-password

# Edit `stack.yml` and replace the image with your own account.
$ faas-cli up
```

Test it out
```bash
curl $OPENFAAS_URL/function/basic-auth --basic --user=admin:password
```

Code: [basic-auth](./basic-auth)

## 2. HMAC

TBH

## 3. Shared API key

TBH


Author: Alex Ellis
