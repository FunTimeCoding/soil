include:
  - orphan

conditional_secret:
  file.managed:
    - name: /etc/conditional
    - contents: secret
