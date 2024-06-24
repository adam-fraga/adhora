package localized

var homePageData = make(map[string]interface{})

func GetHomePageDataEn() map[string]interface{} {
	homePageData["sitename"] = "Astel"
	homePageData["title"] = "Astel Home"
	homePageData["nav_links"] = []string{"Services", "About", "Contact", "Dashboard"}
	homePageData["activity_section_title"] = "Our Activity"
	homePageData["activity_section_content"] = "Astel is a software development company that specializes in building web applications, mobile applications, and websites. We are a team of developers, designers, and project managers that work together to create amazing products for our clients. We are passionate about technology and love what we do. We are always looking for new projects to work on and new clients to work with. If you have a project that you would like us to work on, please contact us. We would love to hear from you."
	homePageData["services_section_title"] = "Services"
	homePageData["services_section_titles"] = []string{
		"Expert Developer Network",
		"Managed Services",
		"Secure Portal Access",
	}
	homePageData["services_section_contents"] = []string{
		"Access a vast network of top-tier developer experts through our web portal. Integrate seamlessly with your team to handle out-of-scope or time-consuming tasks, ensuring high-quality results and innovation.",
		"No need for project management headaches. Our expert-managed services handle time-consuming tasks efficiently, clearing your backlog and allowing your team to focus on core activities.",
		"Enhance productivity with our secure portal access. Communicate effortlessly with developers, schedule meetings, and enjoy a customized environment tailored to your systemâ€™s needs.",
	}
	return homePageData
}

func GetHomePageDataFr() map[string]interface{} {
	homePageData["sitename"] = "Astel"
	homePageData["title"] = "Accueil Astel"
	homePageData["links"] = []string{"Accueil", "Services", "A Propos", "Contact", "Tableau de Bord", "Se Connecter", "S'inscrire"}
	homePageData["activity_section_title"] = "Notre Activité"
	homePageData["activity_section_content"] = "Astel est une entreprise de développement de logiciels spécialisée dans la création d'applications web, d'applications mobiles et de sites web. Nous sommes une équipe de développeurs, de designers et de chefs de projet qui travaillent ensemble pour créer des produits incroyables pour nos clients. Nous sommes passionnés par la technologie et aimons ce que nous faisons. Nous sommes toujours à la recherche de nouveaux projets à réaliser et de nouveaux clients à accompagner. Si vous avez un projet sur lequel vous aimeriez que nous travaillions, veuillez nous contacter. Nous serions ravis de vous entendre."
	homePageData["services_section_title"] = "Services"
	homePageData["services_section_titles"] = []string{
		"Réseau de Développeurs Experts",
		"Services Gérés",
		"Accès au Portail Sécurisé",
	}
	homePageData["services_section_contents"] = []string{
		"Accédez à un vaste réseau d'experts développeurs de premier plan via notre portail web. Intégrez-vous de manière transparente à votre équipe pour gérer les tâches hors champ ou chronophages, garantissant des résultats de haute qualité et de l'innovation.",
		"Plus besoin de maux de tête de gestion de projet. Nos services gérés par des experts traitent efficacement les tâches chronophages, dégageant votre backlog et permettant à votre équipe de se concentrer sur les activités principales.",
		"Améliorez la productivité avec notre accès au portail sécurisé. Communiquez facilement avec les développeurs, planifiez des réunions et profitez d'un environnement personnalisé adapté aux besoins de votre système.",
	}
	return homePageData
}
