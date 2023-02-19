#! /usr/bin/env python
# coding=utf8

import time
from concurrent import futures
import grpc
import sys
import proto_aiservice_pb2_grpc, proto_aiservice_pb2

ONE_DAY_IN_SECONDS = 60 * 60 * 24

PORT = 50002

class AiService(proto_aiservice_pb2_grpc.AiServiceServicer):
    '''
    AiServiceServicer,实现Predict方法
    '''
    def __init__(self):
        pass

    def Predict(self, request, context):
        result = request.image + ":predict done" # 具体算法调用处
        print("call Predict, result=%s", result)
        return proto_aiservice_pb2.PredictResponse(result=result)


def run():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    proto_aiservice_pb2_grpc.add_AiServiceServicer_to_server(AiService(), server)
    server.add_insecure_port('[::]:'+str(PORT))
    server.start()
    print("start service...listen port:", str(PORT))
    try:
        while True:
            time.sleep(ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    run()