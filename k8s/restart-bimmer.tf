resource "kubernetes_cron_job_v1" "bimmerexporter_restart" {
  metadata {
    name = "bimmerexporterrestartt"
  }
  spec {
    concurrency_policy            = "Forbid"
    failed_jobs_history_limit     = 2
    schedule                      = "*/15 * * * *"
    starting_deadline_seconds     = 10
    successful_jobs_history_limit = 2
    job_template {
      metadata {}
      spec {
        backoff_limit              = 2
        ttl_seconds_after_finished = 10
        template {
          metadata {}
          spec {
            service_account_name = "deployment-restart-bimmerexporter"
            container {
              name    = "kubectl"
              image   = "bitnami/kubectl"
              command = ["/bin/sh", "-c","kubectl rollout restart deployment bimmerexporter"]
            }
          }
        }
      }
    }
  }
}