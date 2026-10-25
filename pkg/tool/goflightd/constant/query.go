package constant

const NormalizeTimeQuery = "UPDATE %s SET time = strftime('%%Y-%%m-%%d %%H:%%M:%%S', time) || substr(time, 20, length(time) - 25) || '+00:00' WHERE substr(time, -6) <> '+00:00'"
