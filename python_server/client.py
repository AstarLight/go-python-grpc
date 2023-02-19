#! /usr/bin/env python
# coding=utf8

import grpc
import proto_aiservice_pb2_grpc, proto_aiservice_pb2

TARGET_PORT = 9988

def run():
    conn=grpc.insecure_channel('localhost:'+str(TARGET_PORT))
    client = proto_aiservice_pb2_grpc.GateServiceStub(channel=conn)
    request = proto_aiservice_pb2.ReportRequest(address="127.0.0.1", cpu_load="60%", mem_load="70%")
    response = client.Report(request)
    print("received:",response.status)

if __name__ == '__main__':
    run()