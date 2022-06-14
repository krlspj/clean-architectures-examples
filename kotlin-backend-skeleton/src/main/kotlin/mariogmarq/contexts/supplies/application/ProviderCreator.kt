package mariogmarq.contexts.supplies.application

import mariogmarq.contexts.supplies.domain.City
import mariogmarq.contexts.supplies.domain.Provider
import mariogmarq.contexts.supplies.domain.Repository
import mariogmarq.contexts.supplies.domain.Status

class ProviderCreator(val repository: Repository) {
    fun createProvider(id: Int, name: String, city: City, status: Status) {
        val provider = Provider(id, name, city, status)
        repository.save(provider)
    }
}