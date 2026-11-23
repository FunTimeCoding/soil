absent_secret:
  file.managed:
    - name: /etc/absent
    - contents: secret
