module github.com/krzysbaranski/golang-experiments/hello

go 1.17

replace github.com/krzysbaranski/golang-experiments/greetings => ../greetings

require github.com/krzysbaranski/golang-experiments/greetings v0.0.0-00010101000000-000000000000
