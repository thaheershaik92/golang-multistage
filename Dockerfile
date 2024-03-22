###########################
#Base Image
###########################

From ubuntu As Build

Run apt-get update && apt-get install -y golang-go

Env GO111MODULE=off

Copy . .

Run CGO_ENABLED=0 go build variable.go


##########################
#first stage image
#########################

From scratch

Copy --From=BUILD VARIABLE.GO VARIABLE.GO

ENTRYPOINT ["VARIABLE.GO"]

