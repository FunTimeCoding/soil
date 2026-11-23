include:
  - common

child_package:
  pkg.installed:
    - name: tree
    - require:
      - sls: required
