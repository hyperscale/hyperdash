
// provider_dir = "./providers/"

dashboard {
    title = "My Company Dashboard"

    grid {
        columns = 8
        rows = 6
    }
}

// type, name
tile "status" "api-status" {
    column = 1
    row = 1

    provider "healthcheck" {
        url = "https://api.my-company.cloud/heath"
        timeout = 10
        interval = 10
        unhealthy_threshold = 2
        healthy_threshold = 2
    }
}
