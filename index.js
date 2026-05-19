/**
 * @param params Call parameters, request body for HTTP requests
 * @param context Call context
 *
 * @return Function return data, used as Response Body in HTTP scenarios
 *
 * Full info can be found at:
 * https://inspirecloud.bytedance.net/docs/cloud-function/basic.html
 */
const QRCode = require("qrcode");
var axios = require("axios");
var FormData = require("form-data");
const lark = require("@larksuiteoapi/node-sdk");
let contextLog;
let retryTimes = 0;

const appId = "cli_a021379389b9d013";
const appSecret = "ef9UBv1ElIaaAJSHPrTQCgMguHxUk6l8";
const BaseHost = "gecko.bytedance.net";
const BoeBaseHost = "gecko-boe.bytedance.net";

const APP_NAME_MAP = {
  "1324": "抖音",
  "10281": "抖音",
  "10169": "抖音极速版",
  "10805": "抖音极速版",
  "11309": "抖音火山V2",
  "11617": "抖音火山V2",
  "11401": "生活服务中台",
  "11715": "生活服务中台",
};

const PREFIX_ID_MAP = {
  "1324": "1903",
  "10281": "24",
  "10169": "10073",
  "10805": "280",
  "11309": "10538",
  "11617": "758",
  "11401": "10587",
  "11715": "817"
};

// envType = 'ppe' | 'boe';
let envType = "ppe";
let host = BaseHost;
let ak = '';
let sk = '';

/** Generate QR code base64 */
function getQRCodeBase64(qrCodeScheme) {
  return new Promise((resolve) => {
    QRCode.toDataURL(qrCodeScheme, (err, url) => {
      resolve(url);
    });
  });
}

/** When card sending fails, send error message to the trigger user */
async function sendInitFailMessage(userName, extra) {
  const tenantAccessToken = await getTenantAccessToken();
  const receive_id = await getUserReceiveId(userName);

  const elements = [
    {
      fields: [
        {
          is_short: true,
          text: {
            content: "**失败原因**:\n**没有登记过ak sk，** \n请根据文档提示，登记你的ak sk",
            tag: "lark_md",
          },
        },
        {
          is_short: true,
          text: {
            content:
              "**参考文档:**\nhttps://bytedance.feishu.cn/docx/ChVFdOschojmXmx2L08c8vNlnad\n  「2.登记ak sk」这一段落",
            tag: "lark_md",
          },
        },
      ],
      tag: "div",
    },
    {
      tag: "hr",
    },
    {
      tag: "div",
      text: {
        content: JSON.stringify(extra.params),
        tag: "lark_md",
      },
    },
    {
      tag: "hr",
    },
    {
      elements: [
        {
          content: "来自 gecko-qrcode.js",
          tag: "plain_text",
        },
      ],
      tag: "note",
    },
  ];
  try {
    const res = await axios(
      "https://fsopen.bytedance.net/open-apis/im/v1/messages",
      {
        method: "POST",
        params: {
          receive_id_type: "open_id",
        },
        data: {
          receive_id,
          msg_type: "interactive",
          content: JSON.stringify({
            config: {
              wide_screen_mode: true,
            },
            elements,
            header: {
              template: "red",
              title: {
                content: `【${userName}消息推送失败通知】`,
                tag: "plain_text",
              },
            },
          }),
        },
        headers: {
          Authorization: `Bearer ${tenantAccessToken}`,
        },
      }
    );

    const res2 = await axios(
      "https://fsopen.bytedance.net/open-apis/im/v1/messages",
      {
        method: "POST",
        params: {
          receive_id_type: "chat_id",
        },
        data: {
          receive_id: ChatGroupId,
          msg_type: "interactive",
          content: JSON.stringify({
            config: {
              wide_screen_mode: true,
            },
            elements,
            header: {
              template: "red",
              title: {
                content: `【${userName}消息推送失败通知】`,
                tag: "plain_text",
              },
            },
          }),
        },
        headers: {
          Authorization: `Bearer ${tenantAccessToken}`,
        },
      }
    );

    contextLog("发送日志成功");
    return res;
  } catch (e) {
    contextLog("发送日志失败", e);
  }
}
/** Get receive id */
async function getUserReceiveId(userName) {
  const tenantAccessToken = await getTenantAccessToken();
  const info = await fetch(
    "https://fsopen.bytedance.net/open-apis/contact/v3/users/batch_get_id?user_id_type=open_id",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${tenantAccessToken}`,
      },
      body: JSON.stringify({
        emails: [`${userName}@bytedance.com`],
      }),
    }
  );
  const { data: userInfo } = await info.json();
  const receive_id = userInfo.user_list[0].user_id;
  contextLog("receive_id", receive_id);
  return receive_id;
}

/** When card sending fails, send error message to the trigger user */
async function sendErrorMessage(userName, extra) {
  const tenantAccessToken = await getTenantAccessToken();
  const receive_id = await getUserReceiveId(userName);

  const elements = [
    {
      fields: [
        {
          is_short: true,
          text: {
            content: "**管理员**:\n**@郑南欣** 帮你增加权限",
            tag: "lark_md",
          },
        },
        {
          is_short: true,
          text: {
            content:
              "**参考文档:**\nhttps://bytedance.feishu.cn/docx/ChVFdOschojmXmx2L08c8vNlnad",
            tag: "lark_md",
          },
        },
      ],
      tag: "div",
    },
    {
      tag: "hr",
    },
    {
      tag: "div",
      text: {
        content: JSON.stringify(extra.params),
        tag: "lark_md",
      },
    },
    {
      tag: "hr",
    },
    {
      elements: [
        {
          content: "来自 gecko-qrcode.js",
          tag: "plain_text",
        },
      ],
      tag: "note",
    },
  ];
  try {
    const res = await axios(
      "https://fsopen.bytedance.net/open-apis/im/v1/messages",
      {
        method: "POST",
        params: {
          receive_id_type: "open_id",
        },
        data: {
          receive_id,
          msg_type: "interactive",
          content: JSON.stringify({
            config: {
              wide_screen_mode: true,
            },
            elements,
            header: {
              template: "red",
              title: {
                content: `【${userName}消息推送失败通知】`,
                tag: "plain_text",
              },
            },
          }),
        },
        headers: {
          Authorization: `Bearer ${tenantAccessToken}`,
        },
      }
    );

    const res2 = await axios(
      "https://fsopen.bytedance.net/open-apis/im/v1/messages",
      {
        method: "POST",
        params: {
          receive_id_type: "chat_id",
        },
        data: {
          receive_id: ChatGroupId,
          msg_type: "interactive",
          content: JSON.stringify({
            config: {
              wide_screen_mode: true,
            },
            elements,
            header: {
              template: "red",
              title: {
                content: `【${userName}消息推送失败通知】`,
                tag: "plain_text",
              },
            },
          }),
        },
        headers: {
          Authorization: `Bearer ${tenantAccessToken}`,
        },
      }
    );

    contextLog("发送日志成功");
    return res;
  } catch (e) {
    contextLog("发送日志失败", e);
  }
}

/** Send message */
async function sendCardMessage(trainGroupId, cardData) {
  contextLog("trainGroupId", trainGroupId);
  if (!trainGroupId) {
    return null;
  }

  const tenantAccessToken = await getTenantAccessToken();
  try {
    const res = await axios(
      "https://fsopen.bytedance.net/open-apis/im/v1/messages",
      {
        method: "POST",
        params: {
          receive_id_type: "chat_id",
        },
        data: {
          receive_id: trainGroupId,
          msg_type: "interactive",
          content: JSON.stringify(cardData.card),
        },
        headers: {
          Authorization: `Bearer ${tenantAccessToken}`,
        },
      }
    );

    contextLog("messageData fetch success", res);
    return res;
  } catch (e) {
    contextLog("messageData fetch err", e);
    throw Error(e);
  }
}

/** Self-built app get tenant_access_token */
async function getTenantAccessToken() {
  const res = await axios(
    "https://fsopen.bytedance.net/open-apis/auth/v3/tenant_access_token/internal",
    {
      method: "POST",
      data: {
        app_id: appId,
        app_secret: appSecret,
      },
    }
  );
  return res.data?.tenant_access_token;
}

/** Generate img_key from base64 */
async function getImageKey(base64) {
  try {
    let data = new FormData();
    const base64Data = base64.replace(/^data:image\/\w+;base64,/, "");
    const dataBuffer = new Buffer(base64Data, "base64");
    const tenantAccessToken = await getTenantAccessToken();

    data.append("image_type", "message");
    data.append("image", dataBuffer);

    const config = {
      method: "post",
      url: "https://fsopen.bytedance.net/open-apis/im/v1/images",
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: `Bearer ${tenantAccessToken}`,
        ...data.getHeaders(),
      },
      data: data,
    };

    const res = await axios(config);
    contextLog('请求imageKey res', res);
    return res.data?.data?.image_key;
  } catch (e) {
    contextLog("生成image_key失败", e);
    return;
  }
}

function getSignature(PostBody) {
  // Construct canonicalRequest (final content to be signed)
  const httpMethod = "POST";
  const canonicalURI = "/gecko/server/graphql";
  const canonicalQueryString = "";
  const canonicalHeaders = "content-type;host;x-gecko-user-accesskey";
  const canonicalRequest = `${httpMethod}\n${canonicalURI}\n${canonicalQueryString}\n${canonicalHeaders}\n${PostBody}`;

  const curTimestamp = new Date().getTime();
  const authPrefix = `gecko-auth-v1/${curTimestamp}/30`;
  const signingKeyHmac = crypto.createHmac("sha256", sk);
  signingKeyHmac.update(authPrefix);
  const signingKey = signingKeyHmac.digest("hex");

  // Use signingKey to sign canonicalRequest to get final signature
  const signatureHmac = crypto.createHmac("sha256", signingKey);
  signatureHmac.update(canonicalRequest);
  const signature = signatureHmac.digest("hex");
  return `${authPrefix}/${signature}`;
}

async function getGeckoPackageInfo(package) {
  const { packageID, deploymentId, channel } = package;
  const requestBody = JSON.stringify([
    {
      // operationName: null,
      variables: {
        deploymentId,
        channel: channel,
        targetOS: -1,
        pkgId: String(packageID),
        pkgVersion: "",
        issueStatus: -1,
        issueType: -1,
        cursor: 0,
        count: 20,
      },
      query:
        "query ($deploymentId: Int, $channel: String, $targetOS: Int, $pkgId: String, $pkgVersion: String, $issueStatus: Int, $issueType: Int, $cursor: Int, $count: Int) { packages(deploymentId: $deploymentId, channel: $channel, targetOS: $targetOS, pkgId: $pkgId, pkgVersion: $pkgVersion, issueStatus: $issueStatus, issueType: $issueType, cursor: $cursor, count: $count) { totalCount edges { node { id deploymentId version url creator targetAppVersion targetOS description issueStatus issueType issueValue status channel createdAt updatedAt delIfDownloadFailed delOldPkgBeforeDownload needUnzip needSyncOverseas pkgSize pkgLarkNoticeGroups qrCodeScheme } cursor } pageInfo { startCursor endCursor hasNextPage } } } ",
    },
  ]);

  const graphqlFetchUrl = `https://${host}/gecko/server/graphql`;
  const response = await fetch(graphqlFetchUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "x-gecko-user-accesskey": ak,
      Host: host,
      Authorization: getSignature(requestBody),
    },
    body: requestBody,
  });
  contextLog('getGeckoPackageInfo response', response)
  const [{ data }] = await response.json();
  const { edges = [{}, {}] } = data.packages || {};
  const [{ node = {} }] = edges;

  contextLog("getGeckoPackageInfo edges", edges, data, deploymentId, channel, packageID);
  return node;
}

async function getGeckoChannelId(package) {
  const { packageID, deploymentId, channel } = package;
  const requestBody = JSON.stringify([
    {
      variables: {
        deploymentId,
        name: channel,
        cursor: 0,
        count: 20,
        discard: 0,
        withNotApproved: true,
        channelType: 0,
      },
      query:
        "query ($deploymentId: Int, $name: String, $cursor: Int, $count: Int, $discard: Int, $withNotApproved: Boolean, $channelType: Int) {\n  channels(deploymentId: $deploymentId, name: $name, cursor: $cursor, count: $count, discard: $discard, withNotApproved: $withNotApproved, channelType: $channelType) {\n    totalCount\n    edges {\n      node {\n        id\n        deploymentId\n        name\n        creator\n        modifier\n        larkWebhook\n        larkWebhookName\n        larkWebhookTypes\n        customizedWebhook\n        customizedWebhookTypes\n        customDistributeRule\n        disableDistribute\n        disableDistributeTargetOS\n        disableDistributeTargetAppVersion\n        disablePatch\n        disablePatchTargetOS\n        disablePatchTargetAppVersion\n        enableCDNMultiversion\n        createdAt\n        updatedAt\n        packageType\n        channelType\n        discard\n        serviceTreeId\n        status\n        user {\n          username\n          avatar\n          name\n          email\n        }\n        owners {\n          username\n          avatar\n          name\n          email\n        }\n      }\n      cursor\n    }\n    pageInfo {\n      startCursor\n      endCursor\n      hasNextPage\n    }\n  }\n}\n",
    },
  ]);

  const graphqlFetchUrl = `https://${host}/gecko/server/graphql`;
  const response = await fetch(graphqlFetchUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "x-gecko-user-accesskey": ak,
      Host: host,
      Authorization: getSignature(requestBody),
    },
    body: requestBody,
  });

  const [{ data }] = await response.json();
  const { edges = [{}] } = data.channels || {};
  const [{ node = {} }] = edges;
  let node_id = node.id;
  if(data.channels?.edges && data.channels?.edges.length > 1){
      data.channels?.edges.forEach(item => {
          if(item.node.name == channel){
              node_id = item.node.id;
          }
      })
  }
  contextLog("======node_id", node_id);
  contextLog("======data channel", JSON.stringify(data || {}, null, 2));
  return node_id;
}

/**
 * Get each gecko package info, extract partial info for generating push card content
 */
async function getCardInfo(package) {
  const geckoPackageInfo = await getGeckoPackageInfo(package);
  const geckoChannelId = await getGeckoChannelId(package);

  const base64 = await getQRCodeBase64(geckoPackageInfo?.qrCodeScheme ?? "");
  const imgKey = await getImageKey(base64);
  contextLog("imgKey", imgKey);
  contextLog("geckoInfo", geckoPackageInfo);
  return {
    qrCodeImgKey: imgKey,
    geckoPackageInfo,
    geckoChannelId,
  };
}

/** Get push group id from train config */
function getChatGroupIdList(traingrouplist = []) {
  const FEChatGroupId = "oc_ffe4e047c19114da767fece00f15a138";
  return traingrouplist.filter((i) => i !== FEChatGroupId);
}

/** Get push group id from meego requirement board */
async function getMeegoGroupChatIdList(sdlcinfo) {
  try {
    const meegoInfo = JSON.parse(JSON.parse(sdlcinfo)).workItem_list;
    contextLog("meegoInfo", meegoInfo);

    const pullInAllMeegoGroupPromiseArr = await pullInAllMeegoGroup(meegoInfo);
    const meegoGroupChatIdList = await Promise.all(
      pullInAllMeegoGroupPromiseArr
    );

    return meegoGroupChatIdList
  } catch (error) {
    contextLog('获取meego需求榜单的推送群id 失败', error)
    return []
  }
}

/** 
 * Initialize environment-related variables
 * Check if ak/sk are stored
 */
async function initConfig(envlanename, username) {
  const GeckoAkSkTable = inspirecloud.db.table('gecko_ak_sk');
  const result = await GeckoAkSkTable.where({ userName: username }).findOne();
  if (!result) {
    return false;
  }
  if (envlanename.includes("boe_")) {
    envType = "boe";
    host = BoeBaseHost;
    ak = result.boe_ak;
    sk = result.boe_sk;
  } else {
    envType = "ppe";
    host = BaseHost;
    ak = result.ppe_ak;
    sk = result.ppe_sk;
  }
  if (!ak || !sk) {
    return false
  } else {
    contextLog("initConfig", { username, envType, host, ak, sk });
    return true
  }

}

/** Generate CardData */
async function getLarkCardData(
  projectListParams,
  branch,
  username,
  envlanename
) {
  const promiseArr = projectListParams.map((project) => {
    return getCardInfo(project);
  });

  const cardInfoList = await Promise.all(promiseArr);

  const elements = [];
  let packageChannel = "";

  cardInfoList.forEach((cardInfo) => {
    const {
      qrCodeImgKey,
      geckoPackageInfo: { createdAt, channel, id, deploymentId },
      geckoChannelId,
    } = cardInfo;
    contextLog("cardInfo", cardInfo);

    packageChannel = channel;

    elements.push({
      tag: "column_set",
      flex_mode: "stretch",
      background_style: "grey",
      columns: [
        {
          tag: "column",
          width: "weighted",
          weight: 5,
          vertical_align: "top",
          elements: [
            {
              tag: "markdown",
              content: `**${APP_NAME_MAP[deploymentId]
                }**\n**id**：${id}\n**泳道**：${envlanename}\n[**查看详情**>>](https://cloud${envType === "boe" ? "-boe" : ""
                }.bytedance.net/gecko/site/app/${PREFIX_ID_MAP[deploymentId]
                }/deployment/${deploymentId}/channel/${geckoChannelId}/package)`,
            },
          ],
        },
        {
          tag: "column",
          width: "weighted",
          weight: 1,
          vertical_align: "top",
          elements: [
            {
              tag: "img",
              img_key: `${qrCodeImgKey}`,
              alt: {
                tag: "plain_text",
                content: "",
              },
              mode: "fit_horizontal",
              preview: true,
              compact_width: false,
            },
          ],
        },
      ],
    });
  });

  elements.push({
    tag: "note",
    elements: [
      {
        tag: "plain_text",
        content: `触发人：@${username}    分支：${branch}    channel: ${packageChannel}`,
      },
    ],
  });

  const cardData = {
    msg_type: "interactive",
    card: {
      config: {
        wide_screen_mode: false,
      },
      elements: elements,
      header: {
        template: "turquoise",
        title: {
          content: `🎉 Gecko资源包更新`,
        },
      },
    },
  };

  return cardData;
}

/** Generate meego plugin_token */
async function getPluginToken() {
  const fetchUrl = "https://meego.feishu.cn/open_api/authen/plugin_token";

  const res = await fetch(fetchUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      plugin_id: "MII_63EDEC8584E2401C",
      plugin_secret: "YOUR_LARK_PLUGIN_SECRET",
      state: "111",
      type: 0,
    }),
  });
  contextLog("getPluginToken", res);
  const { data } = await res.json();
  return data.token;
}

/** Pull bot into meego group and get chat_id of the meego requirement group */
async function pullInMeegoGroup(meegoInfo) {
  const { space_key, id, source_type } = meegoInfo;
  contextLog("pullInMeegoGroup", meegoInfo);
  const pluginToken = await getPluginToken();
  const fetchUrl = `https://meego.feishu.cn/open_api/${space_key}/work_item/${id}/bot_join_chat`;

  const res = await fetch(fetchUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-Plugin-Token": pluginToken,
      "X-USER-KEY": "zhengnanxin",
    },
    body: JSON.stringify({
      app_ids: [appId],
      work_item_type_key: source_type ? source_type : 'story',
    }),
  });

  const { data } = await res.json();
  contextLog("pullInMeegoGroup success", data);
  return data.chat_id;
}

/** Automatically pull bot into all meego groups */
async function pullInAllMeegoGroup(meegoInfo = []) {
  return meegoInfo.map((id) => {
    return pullInMeegoGroup(id);
  });
}

/** Send message to all groups */
async function sendAllCardMessage(trainGroupIdList = [], cardData) {
  return trainGroupIdList.map((id) => {
    return sendCardMessage(id, cardData);
  });
}

module.exports = async function run(params, context) {
  try {
    // 1. Set global log printing method
    contextLog = context.log;
    contextLog("params", params, context.headers);

    const { branch, username, traingrouplist = '[]', envlanename, sdlcinfo } =
      context.headers;

    // 2. Initialize config
    const initSuccess = await initConfig(envlanename, username);
    if (!initSuccess) {
      const sendErrorMessageRes = await sendInitFailMessage(username, {
        params,
        headers: context.headers,
        retryTimes,
      });
      return;
    }

    // 3. Get push notification group ids
    const messagePushChatGroupList = getChatGroupIdList(
      JSON.parse(traingrouplist)
    );
    const meegoGroupChatIdList = await getMeegoGroupChatIdList(sdlcinfo);

    // todo: test group, remember to remove after going live
    // const groupIdList = ["oc_739e838a5b0dc4aaf6f43d3c6a6cbd11"];
    const groupIdList = [...messagePushChatGroupList, ...meegoGroupChatIdList];

    if (groupIdList.length > 0) {
      contextLog(
        "groupIdList",
        messagePushChatGroupList,
        meegoGroupChatIdList
      );
    } else {
      // If no groups to notify, terminate directly
      return {
        text: "没有需要推送消息的群"
      };
    }

    // 4. Get package info
    const cardData = await getLarkCardData(
      params,
      branch,
      username,
      envlanename
    );

    // 5. Send package info to meego-bound groups
    const sendMeegoMessagePromiseArr = await sendAllCardMessage(
      [...meegoGroupChatIdList],
      cardData
    );
    const sendRes = await Promise.all(sendMeegoMessagePromiseArr);

    // 6. Send package info to train-bound groups, may not be needed later.
    // Message push failure to train groups is normal, since no one manually added the bot.
    // So for this kind of failure, do not trigger outer retry logic
    try {
      const sendTraninMessagePromiseArr = await sendAllCardMessage(
        [...messagePushChatGroupList, 'oc_739e838a5b0dc4aaf6f43d3c6a6cbd11'],
        cardData
      );
      const sendTrainRes = await Promise.all(sendTraninMessagePromiseArr);
    } catch (e) {
      const last = await sendErrorMessage(context.headers.username, { positon: '消息推送到火车榜单群失败', params, headers: context.headers, e, retryTimes });
    }

  } catch (error) {
    contextLog("sendAllCardMessage err", error);
    const last = await sendErrorMessage(context.headers.username, { params, headers: context.headers, error, retryTimes });

    // Retry once on failure
    if (retryTimes < 1) {
      retryTimes++
      await run(params, context)
    }
  }

  return {
    test: "任务结束",
  };
};
