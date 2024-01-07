<template>
    <AppHeader />
    <main>
        <InputOrganism />
    </main>
</template>

<script lang="ts">
import AppHeader from '@/components/organisms/app_header.vue'
import InputOrganism from '@/components/organisms/input_organism.vue'

export default {
    components: {
        AppHeader,
        InputOrganism
    }
}
</script>
