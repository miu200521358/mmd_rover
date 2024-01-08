import { defineStore } from 'pinia';
import { SaveConfig } from '@/wailsjs/go/config/Config';

export const useLocaleStore = defineStore({
    id: 'locale',
    state: () => ({
        locale: 'ja' // 初期値は日本語
    }),
    actions: {
        setLocale(newLocale: string) {
            console.log('pinia setLocale', this.locale);
            this.locale = newLocale;
            SaveConfig("locale", [newLocale], 1);
            console.log('pinia setLocale new', newLocale);
        }
    }
});
