# select image
FROM mysql

# 環境変数
ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE sample
ENV MYSQL_USER test
ENV MYSQL_PASSWORD password

# 初期実行スクリプト置き場に配置
COPY init/world.sql.gz /docker-entrypoint-initdb.d/world.sql.gz
