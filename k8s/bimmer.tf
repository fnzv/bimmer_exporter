resource "kubernetes_deployment_v1" "bimmerexporter" {
  metadata {
    name = "bimmerexporter"
    labels = {
      test = "bimmerexporter"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        test = "bimmerexporter"
      }
    }

    template {
      metadata {
        labels = {
          test = "bimmerexporter"
        }
      }

      spec {
        container {
          # This image build is public and is the same as the Docker file you find in this repo (if you do not trust it you can build the Dockerfile)
          image = "fnzv/bimmerexporter"
          name  = "bimmerexporter"

          # Those ENV vars are passed externally
          env {
            name  = "EMAIL"
            value = var.bmw_email
          }
          env {
            name  = "PASS"
            value = var.bmw_pass
          }

lifecycle {
            post_start {
                exec {
                    command = [
                        "/bin/bash",
                        "-c",
                        "bimmerconnected fingerprint $EMAIL $PASS rest_of_world && cp /root/vehicle_fingerprint/*/bmw-eadrax-vcs_v4_vehicles_state_WBA0FINGERPRINT01.json /app/bmw.json"
                      ]
                  }
              }
          }
        }
      }
    }
  }
}
resource "kubernetes_ingress_v1" "bimmerexporter" {

   metadata {
       name             = "bimmerexporter"
    }

   spec {

       rule {
           host = "bimmerexporter.youringress.local"

           http {
               path {
                   path      = "/"
                   path_type = "ImplementationSpecific"

                   backend {
                       service {
                           name = "bimmerexporter"

                           port {
                               number = 1337
                            }
                        }
                    }
                }
            }
        }
    }
}

resource "kubernetes_service" "bimmerexporter" {
   metadata {
       name             = "bimmerexporter"
    }

   spec {
       selector                          = {
           "test" = "bimmerexporter"
        }
       type                              = "ClusterIP"

       port {
           port        = 1337
           protocol    = "TCP"
           target_port = "1337"
        }

    }
}



resource "kubernetes_role" "deployment_restart_bimmerexporter" {
  metadata {
    name = "deployment-restart-bimmerexporter"
  }

  rule {
    api_groups = ["apps", "extensions"]
    resources  = ["deployments"]
    resource_names= ["bimmerexporter"]
    verbs      = ["get", "patch", "list", "watch"]
  }
}


resource "kubernetes_service_account" "deployment_restart_bimmerexporter_sa" {
  metadata {
    name = "deployment-restart-bimmerexporter"
  }
}


resource "kubernetes_role_binding_v1" "deployment_restart_bimmerexport_rb2" {
  metadata {
    name      = "deployment-restart-bimmerexporter"
    namespace = "default"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
    name      = "deployment-restart-bimmerexporter"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "deployment-restart-bimmerexporter"
    namespace = "default"
  }
}