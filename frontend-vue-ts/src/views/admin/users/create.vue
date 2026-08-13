<script setup lang="ts">
import {ref, reactive} from 'vue'

import {useRouter} from 'vue-router'

import SidebarMenu from '../../../components/SidebarMenu.vue';

import { useUserCreate } from '../../../composables/user/useUserCreate';

const router = useRouter()

const name = ref<string>('')
const username = ref<string>('')
const email = ref<string>('')
const password = ref<string>('')

const errors = reactive<Record<string, string>>({})

const {mutate, isPending} = useUserCreate()

const storeUser = (e: Event) =>{
    e.preventDefault()

    mutate({
        name: name.value,
        username: username.value,
        email: email.value,
        password: password.value,
    },
    {
        onSuccess: () => {
            router.push('/admin/users')

        },
        onError: (error: any) => {
            Object.assign(errors, error.response.data.errors)
        },
    })
}
</script>

<template>
    <div class="container mt-5 mb-5">
        <div class="row">
            <div class="col-md-3">
                <SidebarMenu />
            </div>
        
        <div class="col-md-9">
            <div class="card border-0 rounded-4 shadow-sm">
                <div class="card-header">
                    ADD USER
                </div>
                <div class="card-body">
                    <form @submit="storeUser">
                        <div class="form-group mb-3">
                            <label class="mb-1 fw-bold">Nama Lengkap</label>
                            <input type="text" v-model="name" class="form-control" placeholder="Nama Lengkap" >
                            <div v-if="errors.Name" class="alert alert-danger mt-2 rounded-4">
                                {{ errors.Name }}
                            </div>
                        </div>

                        <div class="form-group mb-3">
                            <label class="mb-1 fw-bold">Username</label>
                            <input type="text" v-model="username" class="form-control" placeholder="Username">
                            <div v-if="errors.Username" class="alert alert-danger mt-2 rounded-4">
                                {{ errors.Username }}
                            </div>
                        </div>

                         <div class="form-group mb-3">
                            <label class="mb-1 fw-bold">Email</label>
                            <input type="text" v-model="email" class="form-control" placeholder="Email">
                            <div v-if="errors.Email" class="alert alert-danger mt-2 rounded-4">
                                {{ errors.Email }}
                            </div>
                        </div>

                        <div class="form-gorup mb-3">
                            <label class="mb-1 fw-bold">Password</label>
                            <input type="password" v-model="password" class="form-control" placeholder="Password">
                            <div v-if="errors.Password" class="alert alert-danger mt-2 rounded-4">
                                {{ errors.Password }}
                            </div>
                        </div>

                        <button type="submit" class="btn btn-md btn-primary rounded-4 shadow-sm border-0" :disabled="isPending">
                            {{ isPending ? 'Menyimpan...' : 'Simpan' }}
                        </button>

                        <router-link to="/admin/users" class="btn btn-md btn-secondary rounded-4 shadow-sm border-0">Batal</router-link>
                    </form>
                </div>
            </div>
        </div>
    </div>
    </div>
</template>