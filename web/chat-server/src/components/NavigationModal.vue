<template>
  <div class="navigation-bar">
    <div class="up-bar">
      <button class="avatar-btn">
        <el-avatar :src="userInfo.avatar" />
      </button>
    </div>
    <div class="middle-bar">
      <el-tooltip
        effect="customized"
        content="会话聊天"
        placement="left"
        hide-after="0"
        enterable="false"
      >
        <button class="icon-btn" @click="handleToSessionList">
          <el-icon>
            <ChatRound />
          </el-icon>
        </button>
      </el-tooltip>
      <el-tooltip
        effect="customized"
        content="通讯录管理"
        placement="left"
        hide-after="0"
        enterable="false"
      >
        <button class="icon-btn" @click="handleToContactList">
          <el-icon>
            <User />
          </el-icon>
        </button>
      </el-tooltip>
      <el-tooltip
        effect="customized"
        content="朋友圈"
        placement="left"
        hide-after="0"
        enterable="false"
      >
        <button class="icon-btn">
          <el-icon>
            <Share />
          </el-icon>
        </button>
      </el-tooltip>
      <el-tooltip
        effect="customized"
        content="我的收藏"
        placement="left"
        hide-after="0"
        enterable="false"
      >
        <button class="icon-btn">
          <el-icon>
            <Star />
          </el-icon>
        </button>
      </el-tooltip>
      <el-tooltip
        effect="customized"
        content="搜索"
        placement="left"
        hide-after="0"
        enterable="false"
      >
        <button class="icon-btn">
          <el-icon>
            <Search />
          </el-icon>
        </button>
      </el-tooltip>
    </div>
    <div class="down-bar">
      <el-tooltip
        effect="customized"
        content="设置"
        placement="left"
        hide-after="0"
        enterable="false"
      >
        <el-dropdown trigger="click" placement="right">
          <button class="icon-btn">
            <el-icon>
              <Setting />
            </el-icon>
          </button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item
                v-if="userInfo.is_admin == 1"
                @click="handleToManager"
              >
                管理员模式
              </el-dropdown-item>
              <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-tooltip>
      <el-tooltip
        effect="customized"
        content="我的主页"
        placement="left"
        hide-after="0"
        enterable="false"
      >
        <button class="icon-btn" @click="handleToOwnInfo">
          <el-icon><HomeFilled /></el-icon>
        </button>
      </el-tooltip>
    </div>
  </div>
</template>

<script>
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { ElMessage } from "element-plus";
import { reactive, toRefs } from "vue";
import axios from "axios";
export default {
  name: "NavigationModal",
  setup() {
    const router = useRouter();
    const store = useStore();
    const data = reactive({
      userInfo: store.state.userInfo,
    });

    const handleToContactList = () => {
      router.push("/chat/contactlist");
    };

    const handleToSessionList = () => {
      router.push("/chat/sessionlist");
    };

    const handleToManager = () => {
      console.log(data.userInfo);
      router.push("/manager");
    };
    const logout = async () => {
      store.commit("cleanUserInfo");
      const req = {
        owner_id: data.userInfo.uuid,
      };
      const rsp = await axios.post(
        store.state.backendUrl + "/user/wsLogout",
        req
      );
      if (rsp.data.code == 200) {
        router.push("/login");
        ElMessage.success(rsp.data.message);
      } else {
        ElMessage.error(rsp.data.message);
      }
    };
    const handleToOwnInfo = () => {
      router.push("/chat/owninfo");
    };
    return {
      ...toRefs(data),
      router,
      handleToContactList,
      handleToSessionList,
      handleToOwnInfo,
      logout,
      handleToManager,
    };
  },
};
</script>