#!/usr/bin/env python3

import argparse
import json
import logging
import sys

import coloredlogs
from panoramix import decompiler

coloredlogs.install(
    level=logging.INFO,
    fmt="%(asctime)s %(name)s %(message)s",
    datefmt="%H:%M:%S.%f",
    field_styles={"asctime": {"color": "white", "faint": True}},
)

def deco_code(code):
    code = code.strip()
    ctr = decompiler.decompile_bytecode(code)
    d = dict()
    d['asm'] = '\n'.join(ctr.asm)
    d['pseudocode'] = ctr.text
    d['functions'] = ctr.json['functions']
    d['storages'] = ctr.json['stor_defs']
    return d

def deco(args):
    if args.file:
        d = deco_code(args.file.read())
    elif args.stdin:
        d = deco_code(sys.stdin.read())

    doc = json.dumps(d)
    if args.output:
        args.output.write(doc)
    else:
        print(doc)

if __name__ == '__main__':
    parser = argparse.ArgumentParser()

    source = parser.add_mutually_exclusive_group(required=True)
    source.add_argument('-f', '--file', type=argparse.FileType('r'), help='bytecode file in hex')
    source.add_argument('-i', '--stdin', action='store_true', help='read bytecode in hex from stdin')

    parser.add_argument('-o', '--output', type=argparse.FileType('w'))

    args = parser.parse_args()
    deco(args)

