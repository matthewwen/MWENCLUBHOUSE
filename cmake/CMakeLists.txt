cmake_minimum_required(VERSION 3.16)
project(clubhouse)
INCLUDE( FindPkgConfig )

set(CMAKE_C_STANDARD 99)
set(TARGET_NAME ${PROJECT_NAME})
set(CMAKE_SHARED_LINKER_FLAGS "-ljansson -lsqlite3")
#SET(CMAKE_CXX_FLAGS  "${CMAKE_CXX_FLAGS} ${GCC_COVERAGE_COMPILE_FLAGS}")
#SET(CMAKE_EXE_LINKER_FLAGS  "${CMAKE_EXE_LINKER_FLAGS} ${GCC_COVERAGE_LINK_FLAGS}")
#SET(TEMP "${TEMP}${GCC_COVERAGE_COMPILE_FLAGS}" )

include_directories("include")
add_definitions(${CMAKE_SHARED_LINKER_FLAGS})
add_library(clubhouse SHARED lib/apipageview.c include/clubhouse/apipageview.h)
target_link_libraries(clubhouse LINK_PUBLIC ${CMAKE_SHARED_LINKER_FLAGS})

install (TARGETS ${TARGET_NAME}
        RUNTIME DESTINATION bin
        LIBRARY DESTINATION lib
        ARCHIVE DESTINATION lib)

install(DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}/include/clubhouse"
        DESTINATION include
        FILES_MATCHING PATTERN "*.h")


