@REM bat文件名与proto文件名一致
set file_name=%~n0
protoc --proto_path="D:\Python\Python311\Lib\site-packages\grpc_tools\_proto" --proto_path=. --go_out=. --go-grpc_out=. %file_name%.proto