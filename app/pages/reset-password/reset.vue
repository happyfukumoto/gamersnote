<template>
  <div class="signin">
    <BaseHeadline text="パスワード再設定" />
    <div class="email">
      <div class="input">
        <h3>新しいパスワード</h3>
        <BaseInput type="password" :on-input="onInputPassword" />
      </div>
      <BaseButton text="登録" :on-click="resetPassword" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import BaseHeadline from '@/components/BaseHeadline.vue'
import BaseInput from '@/components/BaseInput.vue'
import BaseButton from '@/components/BaseButton.vue'
import { baseModalState } from '@/store'
import { $userApi } from '@/plugins/api'

interface Data {
  password: string
  processing: boolean
  code: string
}

export default Vue.extend({
  components: {
    BaseHeadline,
    BaseInput,
    BaseButton,
  },

  asyncData({ route }): Data | void {
    let code = ''
    if (typeof route.query.code === 'string') {
      code = route.query.code
    }
    return {
      password: '',
      processing: false,
      code,
    }
  },

  data(): Data {
    return {} as Data
  },

  methods: {
    onInputPassword(password: string) {
      this.password = password
    },

    async resetPassword() {
      try {
        this.processing = true
        await $userApi().putPassword({
          password: this.password,
          code: this.code,
        })
        const message = 'パスワードを再設定しました'
        baseModalState.setModal({ showModal: true, message, to: '/' })
      } catch (e) {
        const message = 'パスワードの再設定に失敗しました'
        baseModalState.setModal({ showModal: true, message })
      } finally {
        this.processing = false
      }
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.email {
  padding: 15px 10px;
  max-width: 450px;
  margin: 0 auto;

  h2 {
    color: $dark-grey;
    font-weight: bold;
  }

  .input {
    color: $dark-grey;
    margin: 30px 0;
  }
}
</style>
