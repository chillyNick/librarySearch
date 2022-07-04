# Library Seacrh
This project allows to find books by author or authors by book via grpc request.

## To run project:
1. Run db with test data ``` docker compose up -d ```
2. Run the project ``` make run ```

## Commands from makefile:
- regenerate grpc server ``` make generate-grpc ```
- run tests ``` make test ```