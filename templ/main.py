import re


def main(data):
    pass


def getdata(path):
    with open(path) as f:
        data = f.readlines()
        data = [i.replace('\n', '') for i in data]
    return data


if __name__ == '__main__':
    data = getdata('input.txt')
    testdata = getdata('testinput.txt')
