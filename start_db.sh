docker run --name kasku_db_dev -e MYSQL_ROOT_PASSWORD=K@sKu2023 -e MYSQL_DATABASE=kasku_db -e MYSQL_USER=kasku -e MYSQL_PASSWORD=kasku123 -p 3306:3306 -d mysql:latest --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci