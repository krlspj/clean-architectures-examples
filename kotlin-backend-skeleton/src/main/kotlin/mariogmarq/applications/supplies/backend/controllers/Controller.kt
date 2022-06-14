package mariogmarq.applications.supplies.backend.controllers

import io.ktor.application.*

interface Controller {
    fun run(call: ApplicationCall)
}