package mariogmarq.contexts.supplies.domain

interface Repository {
    fun save(provider: Provider)
    fun getById(id: Int): Provider?
}