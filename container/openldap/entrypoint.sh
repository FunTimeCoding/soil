#!/bin/sh -e

attribute() {
    if [ -z "$3" ]; then
        sed "/$1/d"
    else
        sed "s|$1|$2: $3|"
    fi
}

state=${LDAP_STATE_PATH:-/var/lib/openldap}
template=${LDAP_CONFIG_TEMPLATE:-/etc/openldap/config.ldif.template}
seed=${LDAP_SEED_PATH:-/etc/openldap/seed}
urls=${LDAP_LISTEN_URLS:-ldap:/// ldapi:///}
level=${LDAP_DEBUG_LEVEL:-256}
config=$state/slapd.d
data=$state/data
run=$state/run

if [ -z "$LDAP_SUFFIX" ]; then
    echo "LDAP_SUFFIX is required"
    exit 1
fi

root_dn=${LDAP_ROOT_DN:-cn=admin,$LDAP_SUFFIX}
service_container=${LDAP_SERVICE_CONTAINER:-ou=services}
user_container=${LDAP_USER_CONTAINER:-ou=people}
group_container=${LDAP_GROUP_CONTAINER:-ou=groups}
service_dn=${LDAP_SERVICE_DN:-cn=directory,$service_container,$LDAP_SUFFIX}

mkdir -p "$config" "$data" "$run"

if [ ! -f "$config/cn=config.ldif" ]; then
    if [ -z "$LDAP_ROOT_PASSWORD" ]; then
        echo "LDAP_ROOT_PASSWORD is required to bootstrap $config"
        exit 1
    fi

    hash=$(slappasswd -o module-load=argon2 -h '{ARGON2}' -s "$LDAP_ROOT_PASSWORD")
    sed -e "s|@ROOT_PASSWORD@|$hash|g" \
        -e "s|@ROOT_DN@|$root_dn|g" \
        -e "s|@SUFFIX@|$LDAP_SUFFIX|g" \
        -e "s|@SERVICE_DN@|$service_dn|g" \
        -e "s|@USER_BASE@|$user_container,$LDAP_SUFFIX|g" \
        -e "s|@GROUP_BASE@|$group_container,$LDAP_SUFFIX|g" \
        -e "s|@DATA_PATH@|$data|g" \
        -e "s|@RUN_PATH@|$run|g" \
        "$template" \
        | attribute @TLS_CERTIFICATE@ olcTLSCertificateFile \
            "$LDAP_CERTIFICATE_FILE" \
        | attribute @TLS_KEY@ olcTLSCertificateKeyFile \
            "$LDAP_CERTIFICATE_KEY_FILE" \
        | attribute @TLS_AUTHORITY@ olcTLSCACertificateFile \
            "$LDAP_AUTHORITY_FILE" \
        > /tmp/config.ldif
    slapadd -n0 -F "$config" -l /tmp/config.ldif
    rm /tmp/config.ldif

    head=${LDAP_SUFFIX%%,*}

    if [ "$head" != "${head#dc=}" ]; then
        cat > /tmp/suffix.ldif <<EOF
dn: $LDAP_SUFFIX
objectClass: dcObject
objectClass: organization
dc: ${head#dc=}
o: ${head#dc=}
EOF
        slapadd -F "$config" -l /tmp/suffix.ldif
        rm /tmp/suffix.ldif
    fi

    if [ -n "$LDAP_SERVICE_PASSWORD" ]; then
        service_head=${service_dn%%,*}
        service_hash=$(slappasswd -o module-load=argon2 -h '{ARGON2}' \
            -s "$LDAP_SERVICE_PASSWORD")
        cat > /tmp/service.ldif <<EOF
dn: $service_container,$LDAP_SUFFIX
objectClass: organizationalUnit
ou: ${service_container#ou=}

dn: $service_dn
objectClass: organizationalRole
objectClass: simpleSecurityObject
cn: ${service_head#cn=}
userPassword: $service_hash
EOF
        slapadd -F "$config" -l /tmp/service.ldif
        rm /tmp/service.ldif
    fi

    for file in "$seed"/*.ldif; do
        [ -f "$file" ] || continue
        echo "seeding $file"
        slapadd -F "$config" -l "$file"
    done
fi

chown -R ldap:ldap "$state"

exec slapd -F "$config" -h "$urls" -u ldap -g ldap -d "$level"
