<template>
    <header>
        <LocaleSelect @localeChanged="handleLocaleChange" :initialLocale="initialLocale" />
    </header>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

import LocaleSelect from "@/components/molecules/locale/locale_select.vue";
import { useLocaleStore } from "@/components/molecules/locale/locale_store";

import { SaveConfig, LoadConfig } from "@/wailsjs/go/config/Config";

const initialLocale = ref<string>('ja');
const localeStore = useLocaleStore();

const handleLocaleChange = async (newLocale: string) => {
    console.log("AppHeader: Locale selected:", newLocale);
    await SaveConfig("locale", [newLocale], 1);
    console.log("AppHeader: Locale selected saved:", newLocale);
    localeStore.setLocale(newLocale); // Pinia ストアを更新
};

onMounted(async () => {
    console.log("AppHeader: onMounted");
    const savedLocales = await LoadConfig("locale");
    console.log("AppHeader: onMounted savedLocales:", savedLocales);
    if (savedLocales && savedLocales.length > 0) {
        console.log("AppHeader: Setting locale to", savedLocales[0]);
        initialLocale.value = savedLocales[0]; // 初期Localeを設定
        localeStore.setLocale(savedLocales[0]);
    }
});
</script>
