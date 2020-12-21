openfaas-function-auth
==============

Examples of authentication in OpenFaaS Serverless functions.

## 1. Basic auth

Username password protects governs access for your function. You can also use this with a web-browser, and will receive a prompt to log in.

> Note: you must enable HTTPS / TLS if you access a function with basic auth.

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

The secret and caller share a symmetric key. The caller transmits a hash of the payload, the receiver also computes a hash of the received payload and uses it to verify the result.

See: https://github.com/openfaas/workshop/blob/master/lab11.md

## 3. Shared API key / bearer token

As per #1 without basic auth headers, but using an "Authorization" header.

## 4. OAuth

See the [auth microservice](https://github.com/openfaas/openfaas-cloud/tree/master/auth) in OpenFaaS Cloud which uses OAuth 2.0 with GitHub and GitLab.

## 5. Use a reverse proxy in front of the gateway

Add [Nginx](https://kubernetes.github.io/ingress-nginx/) or similar and implement auth however you like.

See also [ingress-nginx authentication options](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/)

Author: Alex Ellis
