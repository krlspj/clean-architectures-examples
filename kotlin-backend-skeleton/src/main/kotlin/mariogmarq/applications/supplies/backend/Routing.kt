package mariogmarq.applications.supplies.backend

import io.ktor.application.*
import io.ktor.routing.*
import mariogmarq.applications.supplies.backend.controllers.ProviderGetController

fun Application.configureRouting() {
    // Dependencies
    val providerGet = ProviderGetController()

    routing {
        get("/provider/{id}") {
            providerGet.run(call)
        }
    }
}