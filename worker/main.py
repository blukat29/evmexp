#!/usr/bin/env python3

import argparse
import sys
import json

from panoramix import decompiler

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
    print(json.dumps(d))

if __name__ == '__main__':
    parser = argparse.ArgumentParser()

    source = parser.add_mutually_exclusive_group(required=True)
    source.add_argument('-f', '--file', type=argparse.FileType('r'), help='bytecode file in hex')
    source.add_argument('-i', '--stdin', action='store_true', help='read bytecode in hex from stdin')

    args = parser.parse_args()
    deco(args)

