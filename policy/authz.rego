package authz

default allow = false

allow {
    input.method == "GET"
    input.path = ["hello"]
}