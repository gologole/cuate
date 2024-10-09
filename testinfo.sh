#!/bin/bash

# Установите переменные
DB_NAME=db
DB_USER=username
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=5432

docker run --name my_postgres -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=db -p 5432:5432 -d postgres:latest


docker exec -it my_postgres psql -U "$DB_USER" -d postgres -c "CREATE DATABASE $DB_NAME;"


docker exec -it my_postgres psql -U "$DB_USER" -d "$DB_NAME" -c "
CREATE TABLE tasks (
    ig_token TEXT,
    ig_text TEXT,
    ig_photo TEXT,
    vk_access_user_token TEXT,
    vk_access_token_group TEXT,
    vk_group_id TEXT,
    vk_text TEXT,
    vk_photo_path TEXT,
    tg_bot_token TEXT,
    tg_channel_id TEXT,
    tg_text TEXT,
    tg_photo_path TEXT
);"


docker exec -it my_postgres psql -U "$DB_USER" -d "$DB_NAME" -c "
INSERT INTO tasks (
    vk_photo_path,
    tg_photo_path,
    vk_access_user_token,
    vk_access_token_group,
    vk_group_id,
    vk_text,
    tg_bot_token,
    tg_channel_id,
    tg_text
) VALUES
(
    '/home/nikit/Изображения/Снимки экрана/golang.png',
    '/home/nikit/Изображения/Снимки экрана/golang.png',
    'vk1.a.2eETBsSOSwS2_TLU3OYusYjaHUJO2UbaFPsENoHyM9LDf3bbqf-FTxbcIfcV_RC__fkAM14dEAmIi6aporCeFrBsrvQKx7B3ZGKqur2uBc6MUow0_900Ec3KVq0wdmHdCGNWY2xP_XZ963sSpmoROquNw18hHOIDN8lXd3Av7UFKVXVijQvQ_dS79g8yTc7bYYgoZO_h9LRE5vI5fwb15g',
    'vk1.a.JgcEBN0z4uzSTFiMB4S0sUz76SM7MDBEEsqVUCWSab0vxDAZ_jj6eQ2kTv3TKCqKI-7wcJ9McDIAvxysckWFQMNjniKy_spA7vCJ_PaGhLIfbl6qnItvQpZB3l82NL5jOmLxx0hVVgu0s_3dxsaQx_hy3bcTtCbdJ1gLVjOeIHjVCvjbqH_RNUsb4k6ptGTSIteNnXH0U5KnST9hVBdATg',
    '227691717',
    'test post1',
    '7938918286:AAGk2LDDfgOllcVJ1MzhYsJjBy0-bIKR4l4',
    '@testbotschan',
    'test post1'
),
(
    '/home/nikit/Изображения/Снимки экрана/golang.png',
    '/home/nikit/Изображения/Снимки экрана/golang.png',
    'vk1.a.2eETBsSOSwS2_TLU3OYusYjaHUJO2UbaFPsENoHyM9LDf3bbqf-FTxbcIfcV_RC__fkAM14dEAmIi6aporCeFrBsrvQKx7B3ZGKqur2uBc6MUow0_900Ec3KVq0wdmHdCGNWY2xP_XZ963sSpmoROquNw18hHOIDN8lXd3Av7UFKVXVijQvQ_dS79g8yTc7bYYgoZO_h9LRE5vI5fwb15g',
    'vk1.a.JgcEBN0z4uzSTFiMB4S0sUz76SM7MDBEEsqVUCWSab0vxDAZ_jj6eQ2kTv3TKCqKI-7wcJ9McDIAvxysckWFQMNjniKy_spA7vCJ_PaGhLIfbl6qnItvQpZB3l82NL5jOmLxx0hVVgu0s_3dxsaQx_hy3bcTtCbdJ1gLVjOeIHjVCvjbqH_RNUsb4k6ptGTSIteNnXH0U5KnST9hVBdATg',
    '227691717',
    'test post1',
    '7938918286:AAGk2LDDfgOllcVJ1MzhYsJjBy0-bIKR4l4',
    '@testbotschan',
    'test post1'
),
(
    '/home/nikit/Изображения/Снимки экрана/golang.png',
    '/home/nikit/Изображения/Снимки экрана/golang.png',
    'vk1.a.2eETBsSOSwS2_TLU3OYusYjaHUJO2UbaFPsENoHyM9LDf3bbqf-FTxbcIfcV_RC__fkAM14dEAmIi6aporCeFrBsrvQKx7B3ZGKqur2uBc6MUow0_900Ec3KVq0wdmHdCGNWY2xP_XZ963sSpmoROquNw18hHOIDN8lXd3Av7UFKVXVijQvQ_dS79g8yTc7bYYgoZO_h9LRE5vI5fwb15g',
    'vk1.a.JgcEBN0z4uzSTFiMB4S0sUz76SM7MDBEEsqVUCWSab0vxDAZ_jj6eQ2kTv3TKCqKI-7wcJ9McDIAvxysckWFQMNjniKy_spA7vCJ_PaGhLIfbl6qnItvQpZB3l82NL5jOmLxx0hVVgu0s_3dxsaQx_hy3bcTtCbdJ1gLVjOeIHjVCvjbqH_RNUsb4k6ptGTSIteNnXH0U5KnST9hVBdATg',
    '227691717',
    'test post1',
    '7938918286:AAGk2LDDfgOllcVJ1MzhYsJjBy0-bIKR4l4',
    '@testbotschan',
    'test post1'
);"
