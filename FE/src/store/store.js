import api from "@/service/api";
import { createStore } from "vuex";
const storeMahasiswa = {
  state: {
    userMahasiswa: null,
    UserMahasiswaIsLoggedIn: false,
    userMahasiswaRole: null,
    isStoreMahasiswaUpdated: false,
  },
  mutations: {
    setUserMahasiswa(state, payload) {
      state.userMahasiswa = payload;
    },
    setUserMahasiswaIsLoggedIn(state, payload) {
      state.UserMahasiswaIsLoggedIn = payload;
    },
    setUserMahasiswaRole(state, role) {
      state.userMahasiswaRole = role;
    },
    setIsStoreMahasiswaUpdated(state, payload) {
      state.isStoreMahasiswaUpdated = payload;
    },
  },
  actions: {
    async loginMahasiswa(context, { nim, pic }) {
      try {
        const response = await api.LoginPost({
          nim,
          pic,
        });
        document.cookie = `token=${response.data.token}; path=/; secure; HttpOnly`;
        await context.dispatch("updateStoreMahasiswa");
      } catch (err) {
        console.log(err.response.data.error);
        throw new Error(err.response.data.error);
      }
    },
    async updateStoreMahasiswa(context) {
      try {
        const res = await api.getUserData();
        const userData = res.data;
        context.commit("setUserMahasiswa", userData);
        context.commit("setUserMahasiswaIsLoggedIn", true);
        context.commit("setUserMahasiswaRole", userData.user.role);
      } catch (e) {
        console.log(e);
        context.commit("setUserMahasiswa", null);
        context.commit("setUserMahasiswaIsLoggedIn", false);
        context.commit("setUserMahasiswaRole", null);
      }
      context.commit("setIsStoreMahasiswaUpdated", true);
    },
    async logoutMahasiswa(context) {
      try {
        await api.LogoutPost();
        context.commit("setUserMahasiswa", null);
        context.commit("setUserMahasiswaIsLoggedIn", false);
        context.commit("setUserMahasiswaRole", null);
        document.cookie = "token=; path=/; secure; HttpOnly";
      } catch (error) {
        context.commit("setUserMahasiswa", null);
        context.commit("setUserMahasiswaIsLoggedIn", false);
        context.commit("setUserMahasiswaRole", null);
      }
    },
  },
};
const store = createStore({
  modules: {
    storeMahasiswa,
  },
});

export default store;
