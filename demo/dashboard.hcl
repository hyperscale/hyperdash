
// provider_dir = "./providers/"

dashboard {
    title = "My Dashboard"

    server {
        port = 9898
    }

    grid {
        columns = 8
        rows = 6
    }
}

tile "status" "api-status" {
    column = 1
    row = 1

    title = "API Health"

    provider "healthcheck" {
        url = "https://www.google.fr"
        timeout = 5
        interval = 10
        unhealthy_threshold = 2
        healthy_threshold = 2
    }
}
