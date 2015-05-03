net-app-svcs
============

This container image has nginx 1.8.0 with modsecurity 2.9.0 for TLS termination and web application firewall (WAF).

It also has HAProxy 1.5.12 for load balancing.

SSL Key files
-------------

ssl.key and ssl.crt should be replaced with your own key files

To generate your own private key:

    openssl genrsa -des3 -out ssl.key 1024

To remove the passphrase:

    cp ssl.key ssl.key.bak
    openssl rsa -in ssl.key.bak -out ssl.key

To generate a Certificate Signing Request (CSR):

    openssl req -new -key ssl.key -out ssl.csr

To then make a certificate:

    openssl x509 -req -days 3650 -in ssl.csr -signkey ssl.key -out ssl.crt

Config files
------------

haproxy.cfg and nginx.conf can be modified to work with different server IPs etc.
