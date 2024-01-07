<template>
    <div class="text-right mt-2 mr-2 px-10 py-1">
        <select class="bg-gray-500" v-model="selectedLocale" @change="onChangeLocale">
            <option v-for="item in locales" :key="item" :value="item">
                {{ t("locales." + item) }}
            </option>
        </select>
    </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import { useLocaleStore } from './locale_store';

const { t, availableLocales: locales } = useI18n();
const localeStore = useLocaleStore();
const selectedLocale = ref(localeStore.locale);

watch(() => localeStore.locale, (newLocale) => {
    console.log("LocaleSelect: watch localeStore.locale:", newLocale);
    selectedLocale.value = newLocale;
}, { immediate: true });

const onChangeLocale = () => {
    console.log("LocaleSelect: onChangeLocale selectedLocale.value:", selectedLocale.value)
    localeStore.setLocale(selectedLocale.value);
};
</script>
