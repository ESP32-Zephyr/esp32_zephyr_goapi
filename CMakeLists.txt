cmake_minimum_required(VERSION 3.28)
project(esp32_zephyr_goapi)

add_custom_target(go_proto
    COMMAND protoc --proto_path=${CMAKE_CURRENT_SOURCE_DIR}/esp32_zephyr/proto/
    --go_opt=Mcmds.proto=./cmds/
    --go_out=${CMAKE_CURRENT_SOURCE_DIR}
     ${CMAKE_CURRENT_SOURCE_DIR}/esp32_zephyr/proto/cmds.proto
)
