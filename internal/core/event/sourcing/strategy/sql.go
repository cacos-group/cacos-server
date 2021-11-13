package strategy

import "fmt"

const (
	_addInfo           = "INSERT INTO `infos` (`namespace`,`appid`,`level`) VALUES(?,?,?);"
	_createEventLogSql = "CREATE TABLE `%s` (\n  `id` bigint(20) NOT NULL AUTO_INCREMENT,\n  `trx_id` bigint(20) NOT NULL DEFAULT 0,\n  `event_type` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,\n  `args` varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL,\n  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),\n  PRIMARY KEY (`id`),\n  KEY `idx_trx_id` (`trx_id`),\n  KEY `idx_event_type` (`event_type`) \n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_addEventLogSql    = "INSERT INTO `%s` (`event_type`, `args`) VALUES(?,?)"
	_existsEventLogSql = "SELECT count(*) as count FROM `%s` WHERE `event_type` = ? AND `args` = ?"
)

func GenTableName(namespace string, appid string) string {
	return fmt.Sprintf("event_log_%s", namespace)
}