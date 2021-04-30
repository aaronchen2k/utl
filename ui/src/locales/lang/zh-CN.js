import antd from 'ant-design-vue/es/locale-provider/zh_CN'
import momentCN from 'moment/locale/zh-cn'

const components = {
  antLocale: antd,
  momentName: 'zh-cn',
  momentLocale: momentCN
}

const locale = {
  'message': '-',
  'menu.home': '主页',

  'menu.platform': '智能平台',
  'menu.platform.dashboard': '仪表盘',
  'menu.sys.admin': '系统管理',
  'menu.sys.settings': '全局设置',

  'menu.dashboard': '仪表盘',
  'menu.dashboard.analysis': '分析页',
  'menu.dashboard.monitor': '监控页',
  'menu.dashboard.workplace': '工作台',

  'menu.nlu.full': '自然语言理解',
  'menu.nlu': 'NLU',
  'menu.project': '项目',
  'menu.project.list': '项目列表',
  'menu.project.edit': '项目编辑',
  'menu.task': '任务',
  'menu.task.list': '任务列表',
  'menu.task.edit': '任务编辑',
  'menu.intent': '意图',
  'menu.intent.list': '意图列表',
  'menu.intent.edit': '意图编辑',
  'menu.sent': '句子',
  'menu.sent.list': '句子列表',
  'menu.sent.edit': '句子编辑',
  'menu.slot': '语义槽',
  'menu.slot.list': '意语义槽列表',
  'menu.slot.edit': '语义槽编辑',
  'menu.synonym': '同义词',
  'menu.synonym.list': '同义词组',
  'menu.synonym.edit': '编辑词组',
  'menu.synonym.items': '同义词项',
  'menu.synonym.create.item': '创建同义词',
  'menu.synonym.edit.item': '编辑同义词',
  'menu.lookup': '普通词',
  'menu.lookup.list': '普通词组',
  'menu.lookup.edit': '编辑词组',
  'menu.lookup.items': '普通词项',
  'menu.lookup.create.item': '创建普通词',
  'menu.lookup.edit.item': '编辑普通词',

  'form.create': '新建',
  'form.edit': '编辑',
  'form.maintain': '维护',
  'form.remove': '删除',
  'form.disable': '禁用',
  'form.enable': '启用',
  'form.back': '返回',
  'form.save': '保存',
  'form.submit': '提交',
  'form.reset': '重置',
  'form.cancel': '取消',
  'form.search': '查询',
  'form.collapse': '收缩',
  'form.expand': '展开',
  'form.ok': '确认',
  'form.confirm.to.remove': '确认删除？',

  'form.all': '所有',
  'form.no': '编号',
  'form.name': '名称',
  'form.content': '内容',
  'form.status': '状态',
  'form.desc': '描述',
  'form.is.default': '是否默认',
  'form.opt': '操作',

  'status.enable': '启用',
  'status.disable': '禁用',

  'valid.input.name': '请输入名称',
  'valid.input.content': '请输入内容',

  'common.info': '消息',
  'common.tips': '提示',
  'common.confirm': '确认',
  'common.notify': '通知',
  'common.back': '返回',

  'msg.warn': '提醒',
  'msg.confirm.to.logout': '确认退出？',
  'msg.forbidden': '确认退出？',
  'msg.unauthorized': '未授权的',
  'msg.auth.fail': '授权失败',

  'app.setting.pagestyle': '页面演示设置',
  'app.setting.pagestyle.light': '淡色',
  'app.setting.pagestyle.dark': '深色',
  'app.setting.pagestyle.realdark': '深黑',
  'app.setting.themecolor': '主题色彩',
  'app.setting.navigationmode': '导航模式',
  'app.setting.content-width': '内容宽度',
  'app.setting.fixedheader': '固定头部',
  'app.setting.fixedsidebar': '固定菜单栏',
  'app.setting.sidemenu': '左侧菜单栏布局',
  'app.setting.topmenu': '顶部菜单布局',
  'app.setting.content-width.fixed': '固定',
  'app.setting.content-width.fluid': '流式',
  'app.setting.othersettings': '其他设置',
  'app.setting.weakmode': '弱模式',
  'app.setting.copy': '复制设置',
  'app.setting.loading': '加载主题',
  'app.setting.copyinfo': '拷贝成功，请替换src/models/setting.js里的默认设置。',
  'app.setting.production.hint': '设置面板仅显示在开发模式中，请手工进行编辑。'
}

export default {
  ...components,
  ...locale
}
