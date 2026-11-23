required_directory:
  file.directory:
    - name: /srv/required

required_limits:
  file.managed:
    - name: /etc/security/limits.d/required.conf
    - source: salt:///required/limits.conf
