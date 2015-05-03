net-app-svcs
============

This is a version of the haproxy-ssl-ssh container that adds modsecurity and the OWASP rule set to provide a web application firewall (WAF). Also updated HAproxy to 1.5 and added rsyslog.

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
