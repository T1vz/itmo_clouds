terraform {
    required_providers {
        grafana = {
            source = "grafana/grafana"
        }
    }
}

provider "grafana" {
    url = "http://localhost:65239/"
    auth = "glsa_3mw4KBllmDmIAZmB4yWR9yXsGqz3rV5w_e82e322b"
}


resource "grafana_contact_point" "telegram_contact_point" {
  name = "Telegram alerts"

    telegram {
      token    = "6975812664:AAF2cdpF0C5dStHLUfMG6krCkpn_06lQz_A"
      chat_id  = "-4062375569"
    }
}
