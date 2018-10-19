// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Categories", testCategories)
	t.Run("Countries", testCountries)
	t.Run("Currencies", testCurrencies)
	t.Run("Customers", testCustomers)
	t.Run("Descriptions", testDescriptions)
	t.Run("Discounts", testDiscounts)
	t.Run("Manufacturers", testManufacturers)
	t.Run("Materials", testMaterials)
	t.Run("Migrations", testMigrations)
	t.Run("Notifications", testNotifications)
	t.Run("Offers", testOffers)
	t.Run("PasswordResets", testPasswordResets)
	t.Run("Pricelists", testPricelists)
	t.Run("Products", testProducts)
	t.Run("Registries", testRegistries)
	t.Run("RegistryBuilds", testRegistryBuilds)
	t.Run("RegistryFieldStats", testRegistryFieldStats)
	t.Run("RegistryJournals", testRegistryJournals)
	t.Run("RegistryManufacturers", testRegistryManufacturers)
	t.Run("RegistryStatuses", testRegistryStatuses)
	t.Run("Roles", testRoles)
	t.Run("Subcategories", testSubcategories)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Categories", testCategoriesDelete)
	t.Run("Countries", testCountriesDelete)
	t.Run("Currencies", testCurrenciesDelete)
	t.Run("Customers", testCustomersDelete)
	t.Run("Descriptions", testDescriptionsDelete)
	t.Run("Discounts", testDiscountsDelete)
	t.Run("Manufacturers", testManufacturersDelete)
	t.Run("Materials", testMaterialsDelete)
	t.Run("Migrations", testMigrationsDelete)
	t.Run("Notifications", testNotificationsDelete)
	t.Run("Offers", testOffersDelete)
	t.Run("PasswordResets", testPasswordResetsDelete)
	t.Run("Pricelists", testPricelistsDelete)
	t.Run("Products", testProductsDelete)
	t.Run("Registries", testRegistriesDelete)
	t.Run("RegistryBuilds", testRegistryBuildsDelete)
	t.Run("RegistryFieldStats", testRegistryFieldStatsDelete)
	t.Run("RegistryJournals", testRegistryJournalsDelete)
	t.Run("RegistryManufacturers", testRegistryManufacturersDelete)
	t.Run("RegistryStatuses", testRegistryStatusesDelete)
	t.Run("Roles", testRolesDelete)
	t.Run("Subcategories", testSubcategoriesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Categories", testCategoriesQueryDeleteAll)
	t.Run("Countries", testCountriesQueryDeleteAll)
	t.Run("Currencies", testCurrenciesQueryDeleteAll)
	t.Run("Customers", testCustomersQueryDeleteAll)
	t.Run("Descriptions", testDescriptionsQueryDeleteAll)
	t.Run("Discounts", testDiscountsQueryDeleteAll)
	t.Run("Manufacturers", testManufacturersQueryDeleteAll)
	t.Run("Materials", testMaterialsQueryDeleteAll)
	t.Run("Migrations", testMigrationsQueryDeleteAll)
	t.Run("Notifications", testNotificationsQueryDeleteAll)
	t.Run("Offers", testOffersQueryDeleteAll)
	t.Run("PasswordResets", testPasswordResetsQueryDeleteAll)
	t.Run("Pricelists", testPricelistsQueryDeleteAll)
	t.Run("Products", testProductsQueryDeleteAll)
	t.Run("Registries", testRegistriesQueryDeleteAll)
	t.Run("RegistryBuilds", testRegistryBuildsQueryDeleteAll)
	t.Run("RegistryFieldStats", testRegistryFieldStatsQueryDeleteAll)
	t.Run("RegistryJournals", testRegistryJournalsQueryDeleteAll)
	t.Run("RegistryManufacturers", testRegistryManufacturersQueryDeleteAll)
	t.Run("RegistryStatuses", testRegistryStatusesQueryDeleteAll)
	t.Run("Roles", testRolesQueryDeleteAll)
	t.Run("Subcategories", testSubcategoriesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Categories", testCategoriesSliceDeleteAll)
	t.Run("Countries", testCountriesSliceDeleteAll)
	t.Run("Currencies", testCurrenciesSliceDeleteAll)
	t.Run("Customers", testCustomersSliceDeleteAll)
	t.Run("Descriptions", testDescriptionsSliceDeleteAll)
	t.Run("Discounts", testDiscountsSliceDeleteAll)
	t.Run("Manufacturers", testManufacturersSliceDeleteAll)
	t.Run("Materials", testMaterialsSliceDeleteAll)
	t.Run("Migrations", testMigrationsSliceDeleteAll)
	t.Run("Notifications", testNotificationsSliceDeleteAll)
	t.Run("Offers", testOffersSliceDeleteAll)
	t.Run("PasswordResets", testPasswordResetsSliceDeleteAll)
	t.Run("Pricelists", testPricelistsSliceDeleteAll)
	t.Run("Products", testProductsSliceDeleteAll)
	t.Run("Registries", testRegistriesSliceDeleteAll)
	t.Run("RegistryBuilds", testRegistryBuildsSliceDeleteAll)
	t.Run("RegistryFieldStats", testRegistryFieldStatsSliceDeleteAll)
	t.Run("RegistryJournals", testRegistryJournalsSliceDeleteAll)
	t.Run("RegistryManufacturers", testRegistryManufacturersSliceDeleteAll)
	t.Run("RegistryStatuses", testRegistryStatusesSliceDeleteAll)
	t.Run("Roles", testRolesSliceDeleteAll)
	t.Run("Subcategories", testSubcategoriesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Categories", testCategoriesExists)
	t.Run("Countries", testCountriesExists)
	t.Run("Currencies", testCurrenciesExists)
	t.Run("Customers", testCustomersExists)
	t.Run("Descriptions", testDescriptionsExists)
	t.Run("Discounts", testDiscountsExists)
	t.Run("Manufacturers", testManufacturersExists)
	t.Run("Materials", testMaterialsExists)
	t.Run("Migrations", testMigrationsExists)
	t.Run("Notifications", testNotificationsExists)
	t.Run("Offers", testOffersExists)
	t.Run("PasswordResets", testPasswordResetsExists)
	t.Run("Pricelists", testPricelistsExists)
	t.Run("Products", testProductsExists)
	t.Run("Registries", testRegistriesExists)
	t.Run("RegistryBuilds", testRegistryBuildsExists)
	t.Run("RegistryFieldStats", testRegistryFieldStatsExists)
	t.Run("RegistryJournals", testRegistryJournalsExists)
	t.Run("RegistryManufacturers", testRegistryManufacturersExists)
	t.Run("RegistryStatuses", testRegistryStatusesExists)
	t.Run("Roles", testRolesExists)
	t.Run("Subcategories", testSubcategoriesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Categories", testCategoriesFind)
	t.Run("Countries", testCountriesFind)
	t.Run("Currencies", testCurrenciesFind)
	t.Run("Customers", testCustomersFind)
	t.Run("Descriptions", testDescriptionsFind)
	t.Run("Discounts", testDiscountsFind)
	t.Run("Manufacturers", testManufacturersFind)
	t.Run("Materials", testMaterialsFind)
	t.Run("Migrations", testMigrationsFind)
	t.Run("Notifications", testNotificationsFind)
	t.Run("Offers", testOffersFind)
	t.Run("PasswordResets", testPasswordResetsFind)
	t.Run("Pricelists", testPricelistsFind)
	t.Run("Products", testProductsFind)
	t.Run("Registries", testRegistriesFind)
	t.Run("RegistryBuilds", testRegistryBuildsFind)
	t.Run("RegistryFieldStats", testRegistryFieldStatsFind)
	t.Run("RegistryJournals", testRegistryJournalsFind)
	t.Run("RegistryManufacturers", testRegistryManufacturersFind)
	t.Run("RegistryStatuses", testRegistryStatusesFind)
	t.Run("Roles", testRolesFind)
	t.Run("Subcategories", testSubcategoriesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Categories", testCategoriesBind)
	t.Run("Countries", testCountriesBind)
	t.Run("Currencies", testCurrenciesBind)
	t.Run("Customers", testCustomersBind)
	t.Run("Descriptions", testDescriptionsBind)
	t.Run("Discounts", testDiscountsBind)
	t.Run("Manufacturers", testManufacturersBind)
	t.Run("Materials", testMaterialsBind)
	t.Run("Migrations", testMigrationsBind)
	t.Run("Notifications", testNotificationsBind)
	t.Run("Offers", testOffersBind)
	t.Run("PasswordResets", testPasswordResetsBind)
	t.Run("Pricelists", testPricelistsBind)
	t.Run("Products", testProductsBind)
	t.Run("Registries", testRegistriesBind)
	t.Run("RegistryBuilds", testRegistryBuildsBind)
	t.Run("RegistryFieldStats", testRegistryFieldStatsBind)
	t.Run("RegistryJournals", testRegistryJournalsBind)
	t.Run("RegistryManufacturers", testRegistryManufacturersBind)
	t.Run("RegistryStatuses", testRegistryStatusesBind)
	t.Run("Roles", testRolesBind)
	t.Run("Subcategories", testSubcategoriesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Categories", testCategoriesOne)
	t.Run("Countries", testCountriesOne)
	t.Run("Currencies", testCurrenciesOne)
	t.Run("Customers", testCustomersOne)
	t.Run("Descriptions", testDescriptionsOne)
	t.Run("Discounts", testDiscountsOne)
	t.Run("Manufacturers", testManufacturersOne)
	t.Run("Materials", testMaterialsOne)
	t.Run("Migrations", testMigrationsOne)
	t.Run("Notifications", testNotificationsOne)
	t.Run("Offers", testOffersOne)
	t.Run("PasswordResets", testPasswordResetsOne)
	t.Run("Pricelists", testPricelistsOne)
	t.Run("Products", testProductsOne)
	t.Run("Registries", testRegistriesOne)
	t.Run("RegistryBuilds", testRegistryBuildsOne)
	t.Run("RegistryFieldStats", testRegistryFieldStatsOne)
	t.Run("RegistryJournals", testRegistryJournalsOne)
	t.Run("RegistryManufacturers", testRegistryManufacturersOne)
	t.Run("RegistryStatuses", testRegistryStatusesOne)
	t.Run("Roles", testRolesOne)
	t.Run("Subcategories", testSubcategoriesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Categories", testCategoriesAll)
	t.Run("Countries", testCountriesAll)
	t.Run("Currencies", testCurrenciesAll)
	t.Run("Customers", testCustomersAll)
	t.Run("Descriptions", testDescriptionsAll)
	t.Run("Discounts", testDiscountsAll)
	t.Run("Manufacturers", testManufacturersAll)
	t.Run("Materials", testMaterialsAll)
	t.Run("Migrations", testMigrationsAll)
	t.Run("Notifications", testNotificationsAll)
	t.Run("Offers", testOffersAll)
	t.Run("PasswordResets", testPasswordResetsAll)
	t.Run("Pricelists", testPricelistsAll)
	t.Run("Products", testProductsAll)
	t.Run("Registries", testRegistriesAll)
	t.Run("RegistryBuilds", testRegistryBuildsAll)
	t.Run("RegistryFieldStats", testRegistryFieldStatsAll)
	t.Run("RegistryJournals", testRegistryJournalsAll)
	t.Run("RegistryManufacturers", testRegistryManufacturersAll)
	t.Run("RegistryStatuses", testRegistryStatusesAll)
	t.Run("Roles", testRolesAll)
	t.Run("Subcategories", testSubcategoriesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Categories", testCategoriesCount)
	t.Run("Countries", testCountriesCount)
	t.Run("Currencies", testCurrenciesCount)
	t.Run("Customers", testCustomersCount)
	t.Run("Descriptions", testDescriptionsCount)
	t.Run("Discounts", testDiscountsCount)
	t.Run("Manufacturers", testManufacturersCount)
	t.Run("Materials", testMaterialsCount)
	t.Run("Migrations", testMigrationsCount)
	t.Run("Notifications", testNotificationsCount)
	t.Run("Offers", testOffersCount)
	t.Run("PasswordResets", testPasswordResetsCount)
	t.Run("Pricelists", testPricelistsCount)
	t.Run("Products", testProductsCount)
	t.Run("Registries", testRegistriesCount)
	t.Run("RegistryBuilds", testRegistryBuildsCount)
	t.Run("RegistryFieldStats", testRegistryFieldStatsCount)
	t.Run("RegistryJournals", testRegistryJournalsCount)
	t.Run("RegistryManufacturers", testRegistryManufacturersCount)
	t.Run("RegistryStatuses", testRegistryStatusesCount)
	t.Run("Roles", testRolesCount)
	t.Run("Subcategories", testSubcategoriesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Categories", testCategoriesHooks)
	t.Run("Countries", testCountriesHooks)
	t.Run("Currencies", testCurrenciesHooks)
	t.Run("Customers", testCustomersHooks)
	t.Run("Descriptions", testDescriptionsHooks)
	t.Run("Discounts", testDiscountsHooks)
	t.Run("Manufacturers", testManufacturersHooks)
	t.Run("Materials", testMaterialsHooks)
	t.Run("Migrations", testMigrationsHooks)
	t.Run("Notifications", testNotificationsHooks)
	t.Run("Offers", testOffersHooks)
	t.Run("PasswordResets", testPasswordResetsHooks)
	t.Run("Pricelists", testPricelistsHooks)
	t.Run("Products", testProductsHooks)
	t.Run("Registries", testRegistriesHooks)
	t.Run("RegistryBuilds", testRegistryBuildsHooks)
	t.Run("RegistryFieldStats", testRegistryFieldStatsHooks)
	t.Run("RegistryJournals", testRegistryJournalsHooks)
	t.Run("RegistryManufacturers", testRegistryManufacturersHooks)
	t.Run("RegistryStatuses", testRegistryStatusesHooks)
	t.Run("Roles", testRolesHooks)
	t.Run("Subcategories", testSubcategoriesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Categories", testCategoriesInsert)
	t.Run("Categories", testCategoriesInsertWhitelist)
	t.Run("Countries", testCountriesInsert)
	t.Run("Countries", testCountriesInsertWhitelist)
	t.Run("Currencies", testCurrenciesInsert)
	t.Run("Currencies", testCurrenciesInsertWhitelist)
	t.Run("Customers", testCustomersInsert)
	t.Run("Customers", testCustomersInsertWhitelist)
	t.Run("Descriptions", testDescriptionsInsert)
	t.Run("Descriptions", testDescriptionsInsertWhitelist)
	t.Run("Discounts", testDiscountsInsert)
	t.Run("Discounts", testDiscountsInsertWhitelist)
	t.Run("Manufacturers", testManufacturersInsert)
	t.Run("Manufacturers", testManufacturersInsertWhitelist)
	t.Run("Materials", testMaterialsInsert)
	t.Run("Materials", testMaterialsInsertWhitelist)
	t.Run("Migrations", testMigrationsInsert)
	t.Run("Migrations", testMigrationsInsertWhitelist)
	t.Run("Notifications", testNotificationsInsert)
	t.Run("Notifications", testNotificationsInsertWhitelist)
	t.Run("Offers", testOffersInsert)
	t.Run("Offers", testOffersInsertWhitelist)
	t.Run("PasswordResets", testPasswordResetsInsert)
	t.Run("PasswordResets", testPasswordResetsInsertWhitelist)
	t.Run("Pricelists", testPricelistsInsert)
	t.Run("Pricelists", testPricelistsInsertWhitelist)
	t.Run("Products", testProductsInsert)
	t.Run("Products", testProductsInsertWhitelist)
	t.Run("Registries", testRegistriesInsert)
	t.Run("Registries", testRegistriesInsertWhitelist)
	t.Run("RegistryBuilds", testRegistryBuildsInsert)
	t.Run("RegistryBuilds", testRegistryBuildsInsertWhitelist)
	t.Run("RegistryFieldStats", testRegistryFieldStatsInsert)
	t.Run("RegistryFieldStats", testRegistryFieldStatsInsertWhitelist)
	t.Run("RegistryJournals", testRegistryJournalsInsert)
	t.Run("RegistryJournals", testRegistryJournalsInsertWhitelist)
	t.Run("RegistryManufacturers", testRegistryManufacturersInsert)
	t.Run("RegistryManufacturers", testRegistryManufacturersInsertWhitelist)
	t.Run("RegistryStatuses", testRegistryStatusesInsert)
	t.Run("RegistryStatuses", testRegistryStatusesInsertWhitelist)
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
	t.Run("Subcategories", testSubcategoriesInsert)
	t.Run("Subcategories", testSubcategoriesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("DescriptionToProductUsingProduct", testDescriptionToOneProductUsingProduct)
	t.Run("DiscountToCategoryUsingCategory", testDiscountToOneCategoryUsingCategory)
	t.Run("DiscountToManufacturerUsingManufacturer", testDiscountToOneManufacturerUsingManufacturer)
	t.Run("DiscountToSubcategoryUsingSubcategory", testDiscountToOneSubcategoryUsingSubcategory)
	t.Run("ManufacturerToCountryUsingCountry", testManufacturerToOneCountryUsingCountry)
	t.Run("ManufacturerToCurrencyUsingCurrency", testManufacturerToOneCurrencyUsingCurrency)
	t.Run("OfferToCustomerUsingCustomer", testOfferToOneCustomerUsingCustomer)
	t.Run("OfferToProductUsingProduct", testOfferToOneProductUsingProduct)
	t.Run("OfferToUserUsingUser", testOfferToOneUserUsingUser)
	t.Run("PricelistToManufacturerUsingManufacturer", testPricelistToOneManufacturerUsingManufacturer)
	t.Run("PricelistToUserUsingUser", testPricelistToOneUserUsingUser)
	t.Run("ProductToCategoryUsingCategory", testProductToOneCategoryUsingCategory)
	t.Run("ProductToCurrencyUsingCurrency", testProductToOneCurrencyUsingCurrency)
	t.Run("ProductToManufacturerUsingManufacturer", testProductToOneManufacturerUsingManufacturer)
	t.Run("ProductToMaterialUsingMaterial", testProductToOneMaterialUsingMaterial)
	t.Run("ProductToPricelistUsingPricelist", testProductToOnePricelistUsingPricelist)
	t.Run("ProductToSubcategoryUsingSubcategory", testProductToOneSubcategoryUsingSubcategory)
	t.Run("RegistryToRegistryStatusUsingRegistryStatus", testRegistryToOneRegistryStatusUsingRegistryStatus)
	t.Run("RegistryBuildToRegistryUsingRegistry", testRegistryBuildToOneRegistryUsingRegistry)
	t.Run("RegistryFieldStatToRegistryUsingRegistry", testRegistryFieldStatToOneRegistryUsingRegistry)
	t.Run("RegistryFieldStatToRegistryJournalUsingRegistryJournal", testRegistryFieldStatToOneRegistryJournalUsingRegistryJournal)
	t.Run("RegistryManufacturerToRegistryUsingRegistry", testRegistryManufacturerToOneRegistryUsingRegistry)
	t.Run("SubcategoryToCategoryUsingCategory", testSubcategoryToOneCategoryUsingCategory)
	t.Run("UserToRoleUsingRole", testUserToOneRoleUsingRole)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CategoryToDiscounts", testCategoryToManyDiscounts)
	t.Run("CategoryToProducts", testCategoryToManyProducts)
	t.Run("CategoryToSubcategories", testCategoryToManySubcategories)
	t.Run("CountryToManufacturers", testCountryToManyManufacturers)
	t.Run("CurrencyToManufacturers", testCurrencyToManyManufacturers)
	t.Run("CurrencyToProducts", testCurrencyToManyProducts)
	t.Run("CustomerToOffers", testCustomerToManyOffers)
	t.Run("ManufacturerToDiscounts", testManufacturerToManyDiscounts)
	t.Run("ManufacturerToPricelists", testManufacturerToManyPricelists)
	t.Run("ManufacturerToProducts", testManufacturerToManyProducts)
	t.Run("MaterialToProducts", testMaterialToManyProducts)
	t.Run("PricelistToProducts", testPricelistToManyProducts)
	t.Run("ProductToDescriptions", testProductToManyDescriptions)
	t.Run("ProductToOffers", testProductToManyOffers)
	t.Run("RegistryToRegistryBuilds", testRegistryToManyRegistryBuilds)
	t.Run("RegistryToRegistryFieldStats", testRegistryToManyRegistryFieldStats)
	t.Run("RegistryToRegistryManufacturers", testRegistryToManyRegistryManufacturers)
	t.Run("RegistryJournalToRegistryFieldStats", testRegistryJournalToManyRegistryFieldStats)
	t.Run("RegistryStatusToRegistries", testRegistryStatusToManyRegistries)
	t.Run("RoleToUsers", testRoleToManyUsers)
	t.Run("SubcategoryToDiscounts", testSubcategoryToManyDiscounts)
	t.Run("SubcategoryToProducts", testSubcategoryToManyProducts)
	t.Run("UserToOffers", testUserToManyOffers)
	t.Run("UserToPricelists", testUserToManyPricelists)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("DescriptionToProductUsingDescriptions", testDescriptionToOneSetOpProductUsingProduct)
	t.Run("DiscountToCategoryUsingDiscounts", testDiscountToOneSetOpCategoryUsingCategory)
	t.Run("DiscountToManufacturerUsingDiscounts", testDiscountToOneSetOpManufacturerUsingManufacturer)
	t.Run("DiscountToSubcategoryUsingDiscounts", testDiscountToOneSetOpSubcategoryUsingSubcategory)
	t.Run("ManufacturerToCountryUsingManufacturers", testManufacturerToOneSetOpCountryUsingCountry)
	t.Run("ManufacturerToCurrencyUsingManufacturers", testManufacturerToOneSetOpCurrencyUsingCurrency)
	t.Run("OfferToCustomerUsingOffers", testOfferToOneSetOpCustomerUsingCustomer)
	t.Run("OfferToProductUsingOffers", testOfferToOneSetOpProductUsingProduct)
	t.Run("OfferToUserUsingOffers", testOfferToOneSetOpUserUsingUser)
	t.Run("PricelistToManufacturerUsingPricelists", testPricelistToOneSetOpManufacturerUsingManufacturer)
	t.Run("PricelistToUserUsingPricelists", testPricelistToOneSetOpUserUsingUser)
	t.Run("ProductToCategoryUsingProducts", testProductToOneSetOpCategoryUsingCategory)
	t.Run("ProductToCurrencyUsingProducts", testProductToOneSetOpCurrencyUsingCurrency)
	t.Run("ProductToManufacturerUsingProducts", testProductToOneSetOpManufacturerUsingManufacturer)
	t.Run("ProductToMaterialUsingProducts", testProductToOneSetOpMaterialUsingMaterial)
	t.Run("ProductToPricelistUsingProducts", testProductToOneSetOpPricelistUsingPricelist)
	t.Run("ProductToSubcategoryUsingProducts", testProductToOneSetOpSubcategoryUsingSubcategory)
	t.Run("RegistryToRegistryStatusUsingRegistries", testRegistryToOneSetOpRegistryStatusUsingRegistryStatus)
	t.Run("RegistryBuildToRegistryUsingRegistryBuilds", testRegistryBuildToOneSetOpRegistryUsingRegistry)
	t.Run("RegistryFieldStatToRegistryUsingRegistryFieldStats", testRegistryFieldStatToOneSetOpRegistryUsingRegistry)
	t.Run("RegistryFieldStatToRegistryJournalUsingRegistryFieldStats", testRegistryFieldStatToOneSetOpRegistryJournalUsingRegistryJournal)
	t.Run("RegistryManufacturerToRegistryUsingRegistryManufacturers", testRegistryManufacturerToOneSetOpRegistryUsingRegistry)
	t.Run("SubcategoryToCategoryUsingSubcategories", testSubcategoryToOneSetOpCategoryUsingCategory)
	t.Run("UserToRoleUsingUsers", testUserToOneSetOpRoleUsingRole)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("DiscountToSubcategoryUsingDiscounts", testDiscountToOneRemoveOpSubcategoryUsingSubcategory)
	t.Run("ManufacturerToCurrencyUsingManufacturers", testManufacturerToOneRemoveOpCurrencyUsingCurrency)
	t.Run("OfferToCustomerUsingOffers", testOfferToOneRemoveOpCustomerUsingCustomer)
	t.Run("OfferToProductUsingOffers", testOfferToOneRemoveOpProductUsingProduct)
	t.Run("OfferToUserUsingOffers", testOfferToOneRemoveOpUserUsingUser)
	t.Run("ProductToMaterialUsingProducts", testProductToOneRemoveOpMaterialUsingMaterial)
	t.Run("ProductToSubcategoryUsingProducts", testProductToOneRemoveOpSubcategoryUsingSubcategory)
	t.Run("RegistryFieldStatToRegistryUsingRegistryFieldStats", testRegistryFieldStatToOneRemoveOpRegistryUsingRegistry)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CategoryToDiscounts", testCategoryToManyAddOpDiscounts)
	t.Run("CategoryToProducts", testCategoryToManyAddOpProducts)
	t.Run("CategoryToSubcategories", testCategoryToManyAddOpSubcategories)
	t.Run("CountryToManufacturers", testCountryToManyAddOpManufacturers)
	t.Run("CurrencyToManufacturers", testCurrencyToManyAddOpManufacturers)
	t.Run("CurrencyToProducts", testCurrencyToManyAddOpProducts)
	t.Run("CustomerToOffers", testCustomerToManyAddOpOffers)
	t.Run("ManufacturerToDiscounts", testManufacturerToManyAddOpDiscounts)
	t.Run("ManufacturerToPricelists", testManufacturerToManyAddOpPricelists)
	t.Run("ManufacturerToProducts", testManufacturerToManyAddOpProducts)
	t.Run("MaterialToProducts", testMaterialToManyAddOpProducts)
	t.Run("PricelistToProducts", testPricelistToManyAddOpProducts)
	t.Run("ProductToDescriptions", testProductToManyAddOpDescriptions)
	t.Run("ProductToOffers", testProductToManyAddOpOffers)
	t.Run("RegistryToRegistryBuilds", testRegistryToManyAddOpRegistryBuilds)
	t.Run("RegistryToRegistryFieldStats", testRegistryToManyAddOpRegistryFieldStats)
	t.Run("RegistryToRegistryManufacturers", testRegistryToManyAddOpRegistryManufacturers)
	t.Run("RegistryJournalToRegistryFieldStats", testRegistryJournalToManyAddOpRegistryFieldStats)
	t.Run("RegistryStatusToRegistries", testRegistryStatusToManyAddOpRegistries)
	t.Run("RoleToUsers", testRoleToManyAddOpUsers)
	t.Run("SubcategoryToDiscounts", testSubcategoryToManyAddOpDiscounts)
	t.Run("SubcategoryToProducts", testSubcategoryToManyAddOpProducts)
	t.Run("UserToOffers", testUserToManyAddOpOffers)
	t.Run("UserToPricelists", testUserToManyAddOpPricelists)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("CurrencyToManufacturers", testCurrencyToManySetOpManufacturers)
	t.Run("CustomerToOffers", testCustomerToManySetOpOffers)
	t.Run("MaterialToProducts", testMaterialToManySetOpProducts)
	t.Run("ProductToOffers", testProductToManySetOpOffers)
	t.Run("RegistryToRegistryFieldStats", testRegistryToManySetOpRegistryFieldStats)
	t.Run("SubcategoryToDiscounts", testSubcategoryToManySetOpDiscounts)
	t.Run("SubcategoryToProducts", testSubcategoryToManySetOpProducts)
	t.Run("UserToOffers", testUserToManySetOpOffers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("CurrencyToManufacturers", testCurrencyToManyRemoveOpManufacturers)
	t.Run("CustomerToOffers", testCustomerToManyRemoveOpOffers)
	t.Run("MaterialToProducts", testMaterialToManyRemoveOpProducts)
	t.Run("ProductToOffers", testProductToManyRemoveOpOffers)
	t.Run("RegistryToRegistryFieldStats", testRegistryToManyRemoveOpRegistryFieldStats)
	t.Run("SubcategoryToDiscounts", testSubcategoryToManyRemoveOpDiscounts)
	t.Run("SubcategoryToProducts", testSubcategoryToManyRemoveOpProducts)
	t.Run("UserToOffers", testUserToManyRemoveOpOffers)
}

func TestReload(t *testing.T) {
	t.Run("Categories", testCategoriesReload)
	t.Run("Countries", testCountriesReload)
	t.Run("Currencies", testCurrenciesReload)
	t.Run("Customers", testCustomersReload)
	t.Run("Descriptions", testDescriptionsReload)
	t.Run("Discounts", testDiscountsReload)
	t.Run("Manufacturers", testManufacturersReload)
	t.Run("Materials", testMaterialsReload)
	t.Run("Migrations", testMigrationsReload)
	t.Run("Notifications", testNotificationsReload)
	t.Run("Offers", testOffersReload)
	t.Run("PasswordResets", testPasswordResetsReload)
	t.Run("Pricelists", testPricelistsReload)
	t.Run("Products", testProductsReload)
	t.Run("Registries", testRegistriesReload)
	t.Run("RegistryBuilds", testRegistryBuildsReload)
	t.Run("RegistryFieldStats", testRegistryFieldStatsReload)
	t.Run("RegistryJournals", testRegistryJournalsReload)
	t.Run("RegistryManufacturers", testRegistryManufacturersReload)
	t.Run("RegistryStatuses", testRegistryStatusesReload)
	t.Run("Roles", testRolesReload)
	t.Run("Subcategories", testSubcategoriesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Categories", testCategoriesReloadAll)
	t.Run("Countries", testCountriesReloadAll)
	t.Run("Currencies", testCurrenciesReloadAll)
	t.Run("Customers", testCustomersReloadAll)
	t.Run("Descriptions", testDescriptionsReloadAll)
	t.Run("Discounts", testDiscountsReloadAll)
	t.Run("Manufacturers", testManufacturersReloadAll)
	t.Run("Materials", testMaterialsReloadAll)
	t.Run("Migrations", testMigrationsReloadAll)
	t.Run("Notifications", testNotificationsReloadAll)
	t.Run("Offers", testOffersReloadAll)
	t.Run("PasswordResets", testPasswordResetsReloadAll)
	t.Run("Pricelists", testPricelistsReloadAll)
	t.Run("Products", testProductsReloadAll)
	t.Run("Registries", testRegistriesReloadAll)
	t.Run("RegistryBuilds", testRegistryBuildsReloadAll)
	t.Run("RegistryFieldStats", testRegistryFieldStatsReloadAll)
	t.Run("RegistryJournals", testRegistryJournalsReloadAll)
	t.Run("RegistryManufacturers", testRegistryManufacturersReloadAll)
	t.Run("RegistryStatuses", testRegistryStatusesReloadAll)
	t.Run("Roles", testRolesReloadAll)
	t.Run("Subcategories", testSubcategoriesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Categories", testCategoriesSelect)
	t.Run("Countries", testCountriesSelect)
	t.Run("Currencies", testCurrenciesSelect)
	t.Run("Customers", testCustomersSelect)
	t.Run("Descriptions", testDescriptionsSelect)
	t.Run("Discounts", testDiscountsSelect)
	t.Run("Manufacturers", testManufacturersSelect)
	t.Run("Materials", testMaterialsSelect)
	t.Run("Migrations", testMigrationsSelect)
	t.Run("Notifications", testNotificationsSelect)
	t.Run("Offers", testOffersSelect)
	t.Run("PasswordResets", testPasswordResetsSelect)
	t.Run("Pricelists", testPricelistsSelect)
	t.Run("Products", testProductsSelect)
	t.Run("Registries", testRegistriesSelect)
	t.Run("RegistryBuilds", testRegistryBuildsSelect)
	t.Run("RegistryFieldStats", testRegistryFieldStatsSelect)
	t.Run("RegistryJournals", testRegistryJournalsSelect)
	t.Run("RegistryManufacturers", testRegistryManufacturersSelect)
	t.Run("RegistryStatuses", testRegistryStatusesSelect)
	t.Run("Roles", testRolesSelect)
	t.Run("Subcategories", testSubcategoriesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Categories", testCategoriesUpdate)
	t.Run("Countries", testCountriesUpdate)
	t.Run("Currencies", testCurrenciesUpdate)
	t.Run("Customers", testCustomersUpdate)
	t.Run("Descriptions", testDescriptionsUpdate)
	t.Run("Discounts", testDiscountsUpdate)
	t.Run("Manufacturers", testManufacturersUpdate)
	t.Run("Materials", testMaterialsUpdate)
	t.Run("Migrations", testMigrationsUpdate)
	t.Run("Notifications", testNotificationsUpdate)
	t.Run("Offers", testOffersUpdate)
	t.Run("PasswordResets", testPasswordResetsUpdate)
	t.Run("Pricelists", testPricelistsUpdate)
	t.Run("Products", testProductsUpdate)
	t.Run("Registries", testRegistriesUpdate)
	t.Run("RegistryBuilds", testRegistryBuildsUpdate)
	t.Run("RegistryFieldStats", testRegistryFieldStatsUpdate)
	t.Run("RegistryJournals", testRegistryJournalsUpdate)
	t.Run("RegistryManufacturers", testRegistryManufacturersUpdate)
	t.Run("RegistryStatuses", testRegistryStatusesUpdate)
	t.Run("Roles", testRolesUpdate)
	t.Run("Subcategories", testSubcategoriesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Categories", testCategoriesSliceUpdateAll)
	t.Run("Countries", testCountriesSliceUpdateAll)
	t.Run("Currencies", testCurrenciesSliceUpdateAll)
	t.Run("Customers", testCustomersSliceUpdateAll)
	t.Run("Descriptions", testDescriptionsSliceUpdateAll)
	t.Run("Discounts", testDiscountsSliceUpdateAll)
	t.Run("Manufacturers", testManufacturersSliceUpdateAll)
	t.Run("Materials", testMaterialsSliceUpdateAll)
	t.Run("Migrations", testMigrationsSliceUpdateAll)
	t.Run("Notifications", testNotificationsSliceUpdateAll)
	t.Run("Offers", testOffersSliceUpdateAll)
	t.Run("PasswordResets", testPasswordResetsSliceUpdateAll)
	t.Run("Pricelists", testPricelistsSliceUpdateAll)
	t.Run("Products", testProductsSliceUpdateAll)
	t.Run("Registries", testRegistriesSliceUpdateAll)
	t.Run("RegistryBuilds", testRegistryBuildsSliceUpdateAll)
	t.Run("RegistryFieldStats", testRegistryFieldStatsSliceUpdateAll)
	t.Run("RegistryJournals", testRegistryJournalsSliceUpdateAll)
	t.Run("RegistryManufacturers", testRegistryManufacturersSliceUpdateAll)
	t.Run("RegistryStatuses", testRegistryStatusesSliceUpdateAll)
	t.Run("Roles", testRolesSliceUpdateAll)
	t.Run("Subcategories", testSubcategoriesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}