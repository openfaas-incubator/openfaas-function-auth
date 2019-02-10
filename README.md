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

See: https://github.com/openfaas/workshop/blob/master/lab11.md

## 3. Shared API key

As per #1 without basic auth headers, but using an "Authorization" header.

## 4. OAuth

See the [auth microservice](https://github.com/openfaas/openfaas-cloud/tree/master/auth) in OpenFaaS Cloud which uses OAuth 2.0 with GitHub and GitLab.

## 5. Use a reverse proxy in front of the gateway

Add [Kong](https://docs.konghq.com) or similar and implement auth however you like.


Author: Alex Ellis
