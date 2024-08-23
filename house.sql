/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.20.24
 Source Server Type    : MySQL
 Source Server Version : 80024 (8.0.24)
 Source Host           : 192.168.20.24:3306
 Source Schema         : house

 Target Server Type    : MySQL
 Target Server Version : 80024 (8.0.24)
 File Encoding         : 65001

 Date: 06/05/2024 14:32:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 'admin', '431630364@qq.com', 'df9bbce611861b2e5a9a692a8bbe0453');
INSERT INTO `admin` VALUES (2, 'zhangsan', '4244342@qq.com', 'df9bbce611861b2e5a9a692a8bbe0453');

-- ----------------------------
-- Table structure for bulletin
-- ----------------------------
DROP TABLE IF EXISTS `bulletin`;
CREATE TABLE `bulletin`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `date_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Place` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 233 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bulletin
-- ----------------------------
INSERT INTO `bulletin` VALUES (3, '沁心湖放电影啦', '沁心湖民宿将举办露天放映电影活动，邀请您一同来享受户外电影院的别样乐趣！', 'image-4.jpg', '2024-04-30', '沁心湖广场');
INSERT INTO `bulletin` VALUES (232, '泳池派对', '在这个温馨的夜晚，我们诚挚地邀请您加入我们举办的泳池派对。让我们一起在星光闪耀的夜空下，享受温馨的氛围和美好的时刻。', 'b5ed5ec7-84e2-496f-9dd0-24429be235d6.jpg', '2024-04-30T17:37', '沁心湖泳池');

-- ----------------------------
-- Table structure for house
-- ----------------------------
DROP TABLE IF EXISTS `house`;
CREATE TABLE `house`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Price` decimal(10, 2) NULL DEFAULT NULL,
  `house_id` int NULL DEFAULT NULL,
  `Num` int NULL DEFAULT NULL,
  `Description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Area` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Facility` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Policy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `HouseID`(`house_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 107 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of house
-- ----------------------------
INSERT INTO `house` VALUES (1, '山景露台豪华大床房', 1098.00, 1, 3, '整套1室-1床-适用2人', '30', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '7223ae66-39ee-4dd4-93b1-7b74e3dafb40.jpg');
INSERT INTO `house` VALUES (2, '湖景露台豪华大床房', 1098.00, 2, 2, '整套1室-1床-适用2人', '30', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '7.jpg');
INSERT INTO `house` VALUES (3, '庭院竹海家庭套房', 758.00, 3, 1, '整套2室-3床-适用4人', '45', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '13.jpg');
INSERT INTO `house` VALUES (4, '庭院竹海浪漫圆床房', 698.00, 4, 3, '整套1室-1床-适用2人', '30', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '19.jpg');
INSERT INTO `house` VALUES (5, '泳池竹海双床房', 758.00, 5, 2, '整套1室-2床-适用4人', '45', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '25.jpg');
INSERT INTO `house` VALUES (6, '湖景露台豪华双床房', 1098.00, 6, 3, '整套1室-2床-适用4人', '45', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '31.jpg');
INSERT INTO `house` VALUES (7, '双层景观loft', 698.00, 7, 3, '整套1室-2床-适用2人', '50', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '37.jpg');
INSERT INTO `house` VALUES (8, '庭院竹海双床房', 758.00, 8, 1, '整套1室-2床-适用4人', '45', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '43.jpg');
INSERT INTO `house` VALUES (106, '测试房间', 668.00, 0, 10, '', '', '', '', '', 'b20bc954-0a1c-41d7-9f41-2d279eac06c0.jpg');

-- ----------------------------
-- Table structure for houseinfo
-- ----------------------------
DROP TABLE IF EXISTS `houseinfo`;
CREATE TABLE `houseinfo`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Price` decimal(10, 2) NULL DEFAULT NULL,
  `house_id` int NULL DEFAULT NULL,
  `Num` int NULL DEFAULT NULL,
  `Description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Area` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Facility` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Policy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `HouseID`(`house_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 80 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of houseinfo
-- ----------------------------
INSERT INTO `houseinfo` VALUES (1, '山景露台豪华大床房', 998.00, 1, 1, '整套1室-1床-适用2人', '30', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '[\"1.jpg\",\"2.jpg\",\"3.jpg\",\"4.jpg\",\"5.jpg\",\"6.jpg\"]');
INSERT INTO `houseinfo` VALUES (2, '湖景露台豪华大床房', 1098.00, 2, 2, '整套1室-1床-适用2人', '30', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '[\"7.jpg\",\"8.jpg\",\"9.jpg\",\"10.jpg\",\"11.jpg\",\"12.jpg\"]');
INSERT INTO `houseinfo` VALUES (3, '庭院竹海家庭套房', 758.00, 3, 3, '整套2室-3床-适用4人', '45', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '[\"13.jpg\",\"14.jpg\",\"15.jpg\",\"16.jpg\",\"17.jpg\",\"18.jpg\"]');
INSERT INTO `houseinfo` VALUES (4, '庭院竹海浪漫圆床房', 698.00, 4, 3, '整套1室-1床-适用2人', '30', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '[\"19.jpg\",\"20.jpg\",\"21.jpg\",\"22.jpg\",\"23.jpg\",\"24.jpg\"]');
INSERT INTO `houseinfo` VALUES (5, '泳池竹海双床房', 758.00, 5, 3, '整套1室-2床-适用4人', '45', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '[\"25.jpg\",\"26.jpg\",\"27.jpg\",\"28.jpg\",\"29.jpg\",\"30.jpg\"]');
INSERT INTO `houseinfo` VALUES (6, '湖景露台豪华双床房', 1098.00, 6, 2, '整套1室-2床-适用4人', '45', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '[\"31.jpg\",\"32.jpg\",\"33.jpg\",\"34.jpg\",\"35.jpg\",\"36.jpg\"]');
INSERT INTO `houseinfo` VALUES (7, '双层景观loft', 698.00, 7, 3, '整套1室-2床-适用2人', '50', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '[\"37.jpg\",\"38.jpg\",\"39.jpg\",\"40.jpg\",\"41.jpg\",\"42.jpg\"]');
INSERT INTO `houseinfo` VALUES (8, '庭院竹海双床房', 758.00, 8, 1, '整套1室-2床-适用4人', '45', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '[\"43.jpg\",\"44.jpg\",\"45.jpg\",\"46.jpg\",\"47.jpg\",\"48.jpg\"]');
INSERT INTO `houseinfo` VALUES (9, '泳池竹海大床房', 698.00, 9, 3, '整套1室-1床-适用2人', '30', '一张大床（2*1.8米）', '空调，窗户，独立卫浴，无线网络', '入住前两天可免费取消，之后不可取消', '[\"49.jpg\",\"50.jpg\",\"51.jpg\",\"52.jpg\",\"53.jpg\",\"54.jpg\"]');

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `Timestamp` datetime NULL DEFAULT NULL,
  `Message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of messages
-- ----------------------------
INSERT INTO `messages` VALUES (2, '明天会更好', '老板的饭做的很好吃', '2024-04-22 21:24:55', '谢谢您的夸奖');
INSERT INTO `messages` VALUES (3, '朱玉龙', '世界这莫大，我想去看看', '0000-00-00 00:00:00', NULL);
INSERT INTO `messages` VALUES (4, 'user', '11', '0000-00-00 00:00:00', NULL);
INSERT INTO `messages` VALUES (5, 'user', '***', '0000-00-00 00:00:00', NULL);

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Price` decimal(10, 2) NOT NULL,
  `Description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Num` int NULL DEFAULT NULL,
  `Weight` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 226 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES (1, '安徽农家土鸡蛋', 88.00, '来自安徽农家的新鲜牛肉，肉质鲜嫩，口感细腻。', '23820d84-fda9-467b-a9ee-6959f1afec4e.jpg', 20, '2.5kg');
INSERT INTO `products` VALUES (2, '安徽小红肠', 168.00, '精选安徽特色风味原料，制作而成的小红肠，风味独特，回味无穷。', '东北红肠.jpg', 20, '3kg');
INSERT INTO `products` VALUES (3, '安徽油辣椒', 38.00, '采用优质辣椒与特制食用油融合而成，味道鲜香辣口，开胃爽口。', '贵州辣椒.jpg', 20, '0.5kg');
INSERT INTO `products` VALUES (4, '安徽椰粉', 29.00, '选用优质椰子，经过精心加工而成的椰粉，香气扑鼻，口感细腻。', '海南椰粉.jpg', 20, '0.3kg');
INSERT INTO `products` VALUES (5, '黄山烧饼', 45.00, '黄山地区传统特色烧饼，酥香可口，层次分明。', '黄山烧饼.jpg', 20, '20个一箱');
INSERT INTO `products` VALUES (6, '安徽农家李子', 68.00, '来自安徽农家的新鲜李子，酸甜可口，营养丰富。', '李子.jpg', 20, '5kg');
INSERT INTO `products` VALUES (7, '安徽绿茶油', 238.00, '以安徽特产绿茶为原料，精制而成的绿茶油，清香宜人，营养丰富。', '绿茶油.jpg', 20, '1kg');
INSERT INTO `products` VALUES (8, '安徽农家土鸡蛋', 66.00, '农家自养土鸡所产的鸡蛋，营养丰富，味道鲜美。', '农家土鸡蛋.jpg', 20, '8kg');
INSERT INTO `products` VALUES (9, '安徽老陈醋', 23.00, '经过多年陈酿而成的老陈醋，色泽红亮，酸香适口，是安徽地区的传统美食。', '山西陈醋.jpg', 20, '1kg');
INSERT INTO `products` VALUES (10, '安徽农家老蜂蜜', 496.00, '采用农家自养蜂群所产的蜂蜜，纯天然无污染，甜度适中，营养丰富。', '土蜂蜜.jpg', 20, '4kg');
INSERT INTO `products` VALUES (11, '宣城鲜花饼', 66.00, '宣城特色传统饼类，口感香脆，内馅花香四溢，是宣城地区的传统糕点之一。', '鲜花饼.jpg', 20, '30个一箱');
INSERT INTO `products` VALUES (12, '黄山野生菌', 145.00, '产自黄山山脚下的野生菌类，种类丰富，味道鲜美，营养丰富。', '野生菌.jpg', 20, '2kg');
INSERT INTO `products` VALUES (18, '黄山茶油', 336.00, '黄山地区特产茶叶提取的茶油，香气浓郁，口感清爽，营养丰富。', '野生山茶油.jpg', 20, '1kg');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `City` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Age` int NULL DEFAULT NULL,
  `Uid` bigint NOT NULL,
  `Telephone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `CreatedTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdatedTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'user', '62836482@qq.com', 'user', '新乡学院', 24, 348239649, '17837721062', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 'f5230d85-360b-4cf3-8377-dc36e70d8a32.jpg');
INSERT INTO `user` VALUES (2, 'zhangsan', '42340364@qq.com', 'user', '郑州市', 24, 1782440846063308800, '1751611873267', '2024-04-20 21:46:00', '2024-04-24 13:30:06', 'avatar-2.jpg');

-- ----------------------------
-- Table structure for user_house
-- ----------------------------
DROP TABLE IF EXISTS `user_house`;
CREATE TABLE `user_house`  (
  `UUID` bigint NOT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `house_id` bigint NULL DEFAULT NULL,
  `house_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Num` int NULL DEFAULT NULL,
  `Price` decimal(10, 2) NULL DEFAULT NULL,
  `total_price` decimal(10, 2) NULL DEFAULT NULL,
  `created_time` datetime NULL DEFAULT NULL,
  `updated_time` datetime NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`UUID`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_house
-- ----------------------------
INSERT INTO `user_house` VALUES (1783763917227429888, '朱玉龙', 1, '山景露台豪华大床房', '17516118727', '下午3点到\n', 2, 998.00, 1996.00, '2024-04-26 15:43:53', '2024-04-26 15:43:53', 'user');
INSERT INTO `user_house` VALUES (1784403943707643904, '李四', 5, '泳池竹海双床房', '17837721062', '提前帮我把空调打开', 2, 758.00, 1516.00, '2024-04-28 10:07:07', '2024-04-28 10:07:07', '');

-- ----------------------------
-- Table structure for user_product
-- ----------------------------
DROP TABLE IF EXISTS `user_product`;
CREATE TABLE `user_product`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `product_id` int NOT NULL,
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `quantity` int NOT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `telephone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `price` decimal(10, 2) NULL DEFAULT NULL,
  `total_price` decimal(10, 2) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uuid`(`uuid` ASC) USING BTREE,
  INDEX `username`(`username` ASC) USING BTREE,
  INDEX `product_id`(`product_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_product
-- ----------------------------
INSERT INTO `user_product` VALUES (28, '1784967413679263744', 'user', 0, '安徽农家土鸡蛋', 2, '新乡学院D10', '17837721062', 88.00, 176.00, '2024-04-29 23:26:09', '2024-04-30 17:47:04', '', '已发货');
INSERT INTO `user_product` VALUES (29, '1785174204434354176', '李光明', 0, '安徽农家老蜂蜜', 1, '新乡市红旗区大数据产业园', '17837721062', 496.00, 496.00, '2024-04-30 13:07:52', '2024-04-30 17:33:11', '', '未发货');

SET FOREIGN_KEY_CHECKS = 1;
