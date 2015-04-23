#! /bin/gosl

MustSucc(Exec("go", "test"))
MustSucc(Exec("go", "test", "./htmldef"))
