package mariogmarq.applications.supplies.backend.controllers

import io.ktor.application.*
import io.ktor.http.*

class ProviderGetController: Controller {
    override fun run(call: ApplicationCall) {
        call.response.status(HttpStatusCode.OK)
    }
}